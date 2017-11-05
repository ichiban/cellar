package main

import (
	"github.com/goadesign/goa"
	"github.com/ichiban/cellar/app"
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

	// AccountController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	// AccountController_Delete: start_implement

	// Put your logic here

	// AccountController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	// AccountController_List: start_implement

	// Put your logic here

	// AccountController_List: end_implement
	res := app.AccountCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	// AccountController_Show: start_implement

	// Put your logic here

	// AccountController_Show: end_implement
	res := &app.Account{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	// AccountController_Update: start_implement

	// Put your logic here

	// AccountController_Update: end_implement
	return nil
}
