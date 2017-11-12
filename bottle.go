package main

import (
	"github.com/goadesign/goa"

	"github.com/ichiban/cellar/app"
	"github.com/ichiban/cellar/models"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Create runs the create action.
func (c *BottleController) Create(ctx *app.CreateBottleContext) error {
	// BottleController_Create: start_implement

	// Put your logic here
	b := models.Bottle{
		AccountID: ctx.AccountID,
		Name:      ctx.Payload.Name,
		Vineyard:  ctx.Payload.Vineyard,
		Varietal:  ctx.Payload.Varietal,
		Vintage:   ctx.Payload.Vintage,
		Color:     ctx.Payload.Color,
	}
	if err := b.Sweetness.Scan(ctx.Payload.Sweetness); err != nil {
		return err
	}
	if err := b.Country.Scan(ctx.Payload.Country); err != nil {
		return err
	}
	if err := b.Region.Scan(ctx.Payload.Region); err != nil {
		return err
	}
	if err := b.Review.Scan(ctx.Payload.Review); err != nil {
		return err
	}

	// BottleController_Create: end_implement
	return b.Save(DB)
}

// Delete runs the delete action.
func (c *BottleController) Delete(ctx *app.DeleteBottleContext) error {
	// BottleController_Delete: start_implement

	// Put your logic here
	b, err := models.BottleByID(DB, ctx.BottleID)
	if err != nil {
		return err
	}

	// BottleController_Delete: end_implement
	return b.Delete(DB)
}

// List runs the list action.
func (c *BottleController) List(ctx *app.ListBottleContext) error {
	// BottleController_List: start_implement

	// Put your logic here
	bs, err := models.BottlesByAccountID(DB, ctx.AccountID)
	if err != nil {
		return err
	}

	members := make([]*app.HALLink, 0, len(bs))
	for _, b := range bs {
		members = append(members, &app.HALLink{Href: app.BottleHref(b.AccountID, b.ID)})
	}

	// BottleController_List: end_implement
	res := &app.BottleCollection{
		Links: &app.BottleCollectionLinks{
			Self:    &app.HALLink{Href: app.AccountHref(ctx.AccountID) + "/bottles"},
			Members: members,
		},
	}
	return ctx.OK(res)
}

// Rate runs the rate action.
func (c *BottleController) Rate(ctx *app.RateBottleContext) error {
	// BottleController_Rate: start_implement

	// Put your logic here
	b, err := models.BottleByID(DB, ctx.BottleID)
	if err != nil {
		return err
	}
	b.Rating = ctx.Payload.Rating

	// BottleController_Rate: end_implement
	return b.Save(DB)
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// BottleController_Show: start_implement

	// Put your logic here
	b, err := models.BottleByID(DB, ctx.BottleID)
	if err != nil {
		return err
	}

	// BottleController_Show: end_implement
	res := &app.Bottle{
		Links: &app.BottleLinks{
			Self:    &app.HALLink{Href: app.BottleHref(b.AccountID, b.ID)},
			Account: &app.HALLink{Href: app.AccountHref(b.AccountID)},
		},
		ID:       b.ID,
		Name:     b.Name,
		Rating:   &b.Rating,
		Varietal: b.Varietal,
		Vineyard: b.Vineyard,
		Vintage:  b.Vintage,
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *BottleController) Update(ctx *app.UpdateBottleContext) error {
	// BottleController_Update: start_implement

	// Put your logic here
	b, err := models.BottleByID(DB, ctx.BottleID)
	if err != nil {
		return err
	}
	if ctx.Payload.Color != nil {
		b.Color = *ctx.Payload.Color
	}
	if err := b.Country.Scan(ctx.Payload.Country); err != nil {
		return err
	}
	if ctx.Payload.Name != nil {
		b.Name = *ctx.Payload.Name
	}
	if err := b.Region.Scan(ctx.Payload.Region); err != nil {
		return err
	}
	if err := b.Review.Scan(ctx.Payload.Review); err != nil {
		return err
	}
	if err := b.Sweetness.Scan(ctx.Payload.Sweetness); err != nil {
		return err
	}
	if ctx.Payload.Varietal != nil {
		b.Varietal = *ctx.Payload.Varietal
	}
	if ctx.Payload.Vineyard != nil {
		b.Vineyard = *ctx.Payload.Vineyard
	}
	if ctx.Payload.Vintage != nil {
		b.Vintage = *ctx.Payload.Vintage
	}

	// BottleController_Update: end_implement
	return b.Save(DB)
}
