package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
)

func (s *handlerSuite) TestServiceAlertFromModel() {
	model := s.TestData.ServiceAlert_1()
	expected := ServiceAlertV3{
		CommonFields: CommonFields{
			ID: 1,
		},
		Uuid: utils.PointerTo(model.Uuid.String()),
		ServiceAlertV3Create: ServiceAlertV3Create{
			OnEnvironment: utils.PointerTo(s.TestData.Environment_Prod().Name),
			ServiceAlertV3EditableFields: ServiceAlertV3EditableFields{
				Title:        utils.PointerTo(*s.TestData.ServiceAlert_1().Title),
				AlertMessage: utils.PointerTo(*s.TestData.ServiceAlert_1().AlertMessage),
				Link:         utils.PointerTo(*s.TestData.ServiceAlert_1().Link),
				Severity:     utils.PointerTo(*s.TestData.ServiceAlert_1().Severity),
			},
		},
	}
	serviceAlert := ServiceAlertFromModel(model)
	s.Equalf(expected.AlertMessage, serviceAlert.AlertMessage, "ServiceAlertFromModel()")
}

func (s *handlerSuite) TestServiceAlertV3ToModel() {
	ret, err := ServiceAlertV3{
		CommonFields: CommonFields{
			ID: 1,
		},
		Uuid: utils.PointerTo(s.TestData.ServiceAlert_1().Uuid.String()),
		ServiceAlertV3Create: ServiceAlertV3Create{
			OnEnvironment: utils.PointerTo(s.TestData.Environment_Prod().Name),
			ServiceAlertV3EditableFields: ServiceAlertV3EditableFields{
				AlertMessage: utils.PointerTo(*s.TestData.ServiceAlert_1().AlertMessage),
				Title:        utils.PointerTo(*s.TestData.ServiceAlert_1().Title),
				Severity:     utils.PointerTo(*s.TestData.ServiceAlert_1().Severity),
				Link:         utils.PointerTo(*s.TestData.ServiceAlert_1().Link),
			},
		},
	}.toModel(s.DB)
	s.NoError(err)
	if s.NotNil(ret.AlertMessage) {
		s.Equal(*s.TestData.ServiceAlert_1().AlertMessage, *ret.AlertMessage)
	}
}
