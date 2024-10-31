package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/asset_inventory"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/sheets/v4"
	"regexp"
	"strconv"
)

type AssetInventoryV3GenerateRequest struct {
	GoogleSheetURL  string `json:"googleSheetURL"`
	SkipLeadingRows int    `json:"skipLeadingRows" default:"1"`
}

type AssetInventoryV3GenerateResponse struct {
}

// assetInventoryProceduresV3Generate godoc
//
//	@summary		Generate an asset inventory
//	@description	Update the given Google Sheet with the latest asset inventory data
//	@tags			AssetInventory
//	@accept			json
//	@produce		json
//	@param			config					body		AssetInventoryV3GenerateRequest	true	"Configuration for the request"
//	@success		200						{object}	AssetInventoryV3GenerateResponse
//	@failure		400,403,404,407,409,500	{object}	errors.ErrorResponse
//	@router			/api/asset-inventory/procedures/v3/generate [post]
func assetInventoryProceduresV3Generate(ctx *gin.Context) {
	// Authenticate request
	user, err := authentication.MustUseUser(ctx)
	if err != nil {
		return
	}
	db, err := authentication.MustUseDB(ctx)
	if err != nil {
		return
	}

	// Bail if this endpoint is disabled in config
	if !config.Config.Bool("assetInventory.enable") {
		errors.AbortRequest(ctx, fmt.Errorf("asset inventory functionality is disabled"))
		return
	}

	// Authorize request based on config
	if requiredRoleReference := config.Config.String("assetInventory.requiredRole"); requiredRoleReference != "" {
		var requiredRole, requiredRoleQuery models.Role
		if requiredRoleQuery, err = roleModelFromSelector(requiredRoleReference); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error parsing assetInventory.requiredRole: %w", err))
			return
		} else if err = db.Where(&requiredRoleQuery).Select("id").First(&requiredRole).Error; err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("error querying assetInventory.requiredRole %s: %w", config.Config.String("assetInventory.requiredRole"), err))
			return
		} else if err = user.ErrIfNotActiveInRole(db, &requiredRole.ID); err != nil {
			errors.AbortRequest(ctx, fmt.Errorf("caller not authorized to generate asset inventory: %w", err))
			return
		}
	}

	// Parse request
	var body AssetInventoryV3GenerateRequest
	if err = ctx.ShouldBindJSON(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) request validation error: %w", errors.BadRequest, err))
		return
	}
	if err = defaults.Set(&body); err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error setting defaults: %w", err))
		return
	}

	// Parse Google Sheets URL
	//   - Format per https://developers.google.com/sheets/api/guides/concepts
	//   - https://docs.google.com/spreadsheets/d/<SPREADSHEET_ID>/edit?gid=<SHEET_ID>#gid=<SHEET_ID>
	var spreadsheetID, sheetID string
	if matches := regexp.
		MustCompile(`https://docs.google.com/spreadsheets/d/([^/]+)/`).
		FindStringSubmatch(body.GoogleSheetURL); len(matches) != 2 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) invalid Google Sheet URL", errors.BadRequest))
		return
	} else {
		spreadsheetID = matches[1]
	}
	if matches := regexp.
		MustCompile(`[#&?]gid=([^#&?]+)`).
		FindStringSubmatch(body.GoogleSheetURL); len(matches) != 2 {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) invalid Google Sheet URL", errors.BadRequest))
		return
	} else {
		sheetID = matches[1]
	}

	// Load spreadsheet
	sheetsService, err := sheets.NewService(ctx)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error creating Google Sheets service: %w", err))
		return
	}
	spreadsheetData, err := sheetsService.Spreadsheets.Get(spreadsheetID).Context(ctx).Do()
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error querying Google Sheets service for spreadsheet metadata: %w", err))
		return
	}
	var sheetTitle string
	for _, sheet := range spreadsheetData.Sheets {
		if sheet.Properties != nil && strconv.FormatInt(sheet.Properties.SheetId, 10) == sheetID {
			sheetTitle = sheet.Properties.Title
			break
		}
	}
	if sheetTitle == "" {
		errors.AbortRequest(ctx, fmt.Errorf("(%s) Google Sheet URL possibly incorrect, no match for sheet ID found", errors.BadRequest))
		return
	}
	existingSpreadsheetValues, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, sheetTitle).Context(ctx).Do()
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error querying Google Sheets service: %w", err))
		return
	}

	spreadsheet, err := asset_inventory.NewSpreadsheet(existingSpreadsheetValues.Values, body.SkipLeadingRows)
	if err != nil {
		errors.AbortRequest(ctx, fmt.Errorf("error creating asset inventory spreadsheet object: %w", err))
		return
	}

	// grab new data from the cloud asset api

	// update the spreadsheet with the new data

	// This call won't print much because there's nothing that reads in the current assets to it... it's just sorta a demo
	fmt.Printf("Values: %v\n", spreadsheet.ToRows())
}
