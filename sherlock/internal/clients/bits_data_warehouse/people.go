package bits_data_warehouse

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/iterator"
)

type Person struct {
	// ManagerUsername is the Username of the Person's manager according to Broad HR systems.
	// This may not be who the person actually reports to, especially in the case of partner
	// organizations. Folks from other organizations will often have one of our administrators
	// listed as their "manager" in this field.
	ManagerUsername bigquery.NullString `bigquery:"manager_username"`
	// Username is generally the email prefix of the Person's BroadEmail.
	Username bigquery.NullString `bigquery:"username"`
	// Phones contains any Broad-assigned phone number for the Person. When this is present,
	// it tends to follow the format of "+d (ddd) ddd-dddd" where "d" is a digit.
	Phones bigquery.NullString `bigquery:"phones"`
	// Locations contains any Broad-assigned desk or office for the Person. Due to hybrid
	// work policies, this often references a "home base" rather than an actual working
	// location.
	Locations bigquery.NullString `bigquery:"locations"`
	// Title is the Person's "business title" according to Broad HR systems. This is what's
	// shown as the Person's title in Slack or other systems, but it is separate from the
	// person's "position" or "job profile" in Workday which is more standardized.
	Title bigquery.NullString `bigquery:"title"`
	// FirstName is according to Broad's HR systems.
	FirstName bigquery.NullString `bigquery:"first_name"`
	// LastName is according to Broad's HR systems.
	LastName bigquery.NullString `bigquery:"last_name"`
	// Group is the org within the Broad where the Person works. This field isn't always
	// updated during reorgs or job changes so it can be out of date.
	Group bigquery.NullString `bigquery:"group"`
	// Classification indicates the sort of employee the Person is. For our org this
	// mainly just indicates whether someone is on Broad's payroll directly or whether
	// they're just Broad-badged. "Employee" is the quote-on-quote normal value here.
	Classification bigquery.NullString `bigquery:"classification"`
	// Affiliation is the Person's home organization. For folks with a Classification
	// of "Employee" this will be the Broad, but for merely Broad-badged folks it'll
	// indicate what organization they're actually employed by.
	Affiliation bigquery.NullString `bigquery:"affiliation"`
	// BroadEmail is the Person's Broad email, which should hopefully match with their
	// GoogleUsername (but might not, because it's possible to change stuff in Google
	// Workspace).
	BroadEmail bigquery.NullString `bigquery:"broad_email"`
	// GoogleUsername is how the Person will be identified to us by Google. This is the
	// value that we can most-safely match to the models.User.Email field we store,
	// because we define that field based on what we receive from Google's IAP.
	GoogleUsername bigquery.NullString `bigquery:"google_username"`
}

var (
	cachedPeople   map[string]Person
	cacheUpdatedAt time.Time
	personColumns  string
)

// calculatePersonColumns assembles a list of the columns from the struct that we'll
// scan them into. This is a dumb thing but we'll just run it once and it means we
// don't have to reference the columns twice.
func calculatePersonColumns() {
	typ := reflect.TypeFor[Person]()
	for i := 0; i < typ.NumField(); i++ {
		if i > 0 {
			personColumns += ", "
		}
		columnName, _ := strings.CutSuffix(typ.Field(i).Tag.Get("bigquery"), ",")
		personColumns += "`" + columnName + "`" // Escape with backticks to avoid error on "group" column
	}
}

func updatePeopleCache(ctx context.Context) error {
	queryString := "SELECT " + personColumns + " FROM " + config.Config.String("bitsDataWarehouse.peopleTable")
	query := client.Query(queryString)
	it, err := query.Read(ctx)
	if err != nil {
		return err
	}
	newCache := make(map[string]Person)
	for {
		var p Person
		err = it.Next(&p)
		if errors.Is(err, iterator.Done) {
			break
		} else if err != nil {
			return err
		} else if p.GoogleUsername.Valid {
			newCache[p.GoogleUsername.String()] = p
		}
	}
	cachedPeople = newCache
	cacheUpdatedAt = time.Now()
	return nil
}

func keepPeopleCacheUpdated(ctx context.Context) {
	interval := config.Config.MustDuration("bitsDataWarehouse.peopleUpdateInterval")
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
			if err := updatePeopleCache(ctx); err != nil {
				log.Warn().Err(err).Msgf("failed to update people cache, last updated at %s: %v", cacheUpdatedAt.String(), err)
			}
		}
	}
}

// GetPerson retrieves a Person from the cache by email address. Since it's cached, there's no monetary cost
// associated with calling this function. If the cache hasn't been initialized, an error will be returned.
// The caller should handle this as the person being neither found nor not found.
//
// We disable linting here because we're just building functionality for later; we don't have an exact use
// case we're shipping right now.
//
//nolint:unused
//goland:noinspection GoUnusedExportedFunction,GoUnnecessarilyExportedIdentifiers
func GetPerson(email string) (person Person, found bool, err error) {
	if len(cachedPeople) == 0 {
		return Person{}, false, errors.New("people cache not initialized")
	} else {
		person, found = cachedPeople[email]
		return person, found, nil
	}
}

func UseMockedGetPerson(t *testing.T, mockPeople map[string]Person, callback func()) {
	if t == nil {
		panic("UseMockedGetPerson must be called with a non-nil testing.T")
	}
	existingCache := cachedPeople
	cachedPeople = mockPeople
	callback()
	cachedPeople = existingCache
}
