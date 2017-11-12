package main

import (
	"github.com/goadesign/goa"
	"github.com/ichiban/cellar/app"
	"github.com/ichiban/cellar/models"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController")}
}

// Create runs the create action.
func (c *AccountController) Create(ctx *app.CreateAccountContext) error {
	// AccountController_Create: start_implement

	// Put your logic here
	a := models.Account{
		Name: ctx.Payload.Name,
	}

	// AccountController_Create: end_implement
	return a.Save(DB)
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	// AccountController_Delete: start_implement

	// Put your logic here
	a, err := models.AccountByID(DB, ctx.AccountID)
	if err != nil {
		return err
	}

	// AccountController_Delete: end_implement
	return a.Delete(DB)
}

// List runs the list action.
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	// AccountController_List: start_implement

	// Put your logic here
	as, err := models.GetAccountIDs(DB)
	if err != nil {
		return err
	}

	var members []*app.HALLink
	for _, a := range as {
		members = append(members, &app.HALLink{Href: app.AccountHref(a.ID)})
	}

	// AccountController_List: end_implement
	res := &app.AccountCollection{
		Links: &app.AccountCollectionLinks{
			Self:    &app.HALLink{Href: "/accounts"},
			Members: members,
		},
	}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	// AccountController_Show: start_implement

	// Put your logic here
	a, err := models.AccountByID(DB, ctx.AccountID)
	if err != nil {
		return err
	}

	// AccountController_Show: end_implement
	res := &app.Account{
		Links: &app.AccountLinks{
			Self:    &app.HALLink{Href: app.AccountHref(a.ID)},
			Bottles: &app.HALLink{Href: app.AccountHref(a.ID) + "/bottles"},
		},
		ID:        a.ID,
		Name:      a.Name,
		CreatedAt: a.CreatedAt.Time,
		CreatedBy: a.CreatedBy,
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	// AccountController_Update: start_implement

	// Put your logic here
	a, err := models.AccountByID(DB, ctx.AccountID)
	if err != nil {
		return err
	}

	// AccountController_Update: end_implement
	return a.Delete(DB)
}
