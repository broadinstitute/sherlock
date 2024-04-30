package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

// environmentsV3Upsert godoc
//
//	@summary		Upsert a dynamic Environment
//	@description	Create or get a dynamic environment, depending on whether it already exists. If it already exists, you'll be marked as its owner. It refuses to work with non-dynamic environments (you can't specify a lifecycle of "static" or "template").
//	@tags			Environments
//	@accept			json
//	@produce		json
//	@param			environment				body		EnvironmentV3Create	true	"The Environment to upsert"
//	@success		201						{object}	EnvironmentV3
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/environments/v3 [put]
func environmentsV3Upsert(ctx *gin.Context) {
	user, err := authentication.MustUseUser(ctx)
	if err != nil {
		return
	}

	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	var body EnvironmentV3Create
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}

	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults for environment: %w", err))
		return
	}

	// Even if the user didn't specify "dynamic", we would've just filled it in ourselves from
	// defaults.Set(&body) above. If the user specified "static" or "template", we definitely
	// don't want them using upsert, so we bail out.
	if body.Lifecycle != "dynamic" {
		errors.AbortRequest(ctx, fmt.Errorf("lifecycle for upsert must be dynamic"))
		return
	}

	toUpsert, err := body.toModel(db)
	if err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	upsertTransaction := db.Where(&toUpsert)

	// If we're not upserting a specific BEE, only match ones that aren't allocated to someone.
	//
	// 		Fun fact: Sherlock happens to enforce against null owners right now, so this'll always
	// 		fail to find a match, just for the moment. That's okay!
	//
	// 		When we add pooling, I'm making a guess that we'll implement that by making owner a nullable
	// 		field, and unallocated BEEs will just have no owner yet. Maybe we'll make a separate field
	// 		to represent an unallocated BEE... if we do, we should update this query to match that field
	// 		instead.
	//
	// 		What does this query do right now since owner can't be null? It means we'll always fail to
	// 		find a match, so we'll always create a new BEE if a name isn't passed.
	if toUpsert.Name == "" {
		upsertTransaction = upsertTransaction.Where("owner_id IS NULL AND legacy_owner IS NULL")
	}

	// Regardless of whether we're adding a BEE or finding one, we always want to claim ownership of it.
	upsertTransaction.Assign(&models.Environment{OwnerID: &user.ID})

	// Run the upsert transaction
	var result models.Environment
	upsertTransaction.FirstOrCreate(&result)
	if err = upsertTransaction.Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	// Now we load the environment, with all the associations, to return.
	//
	// 		We can reuse the same variable here because we're done with the upsert. The reason we don't
	// 		add a preload clause to that upsert is because it is potentially a mutation operation, and
	// 		we never want to have a preload and a mutation coexist (it can update the associations!)
	if err = db.Preload(clause.Associations).First(&result, result.ID).Error; err != nil {
		errors.AbortRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, environmentFromModel(result))
}
