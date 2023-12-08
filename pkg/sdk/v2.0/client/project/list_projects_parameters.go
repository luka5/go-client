// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListProjectsParams creates a new ListProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectsParams() *ListProjectsParams {
	return &ListProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectsParamsWithTimeout creates a new ListProjectsParams object
// with the ability to set a timeout on a request.
func NewListProjectsParamsWithTimeout(timeout time.Duration) *ListProjectsParams {
	return &ListProjectsParams{
		timeout: timeout,
	}
}

// NewListProjectsParamsWithContext creates a new ListProjectsParams object
// with the ability to set a context for a request.
func NewListProjectsParamsWithContext(ctx context.Context) *ListProjectsParams {
	return &ListProjectsParams{
		Context: ctx,
	}
}

// NewListProjectsParamsWithHTTPClient creates a new ListProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectsParamsWithHTTPClient(client *http.Client) *ListProjectsParams {
	return &ListProjectsParams{
		HTTPClient: client,
	}
}

/*
ListProjectsParams contains all the parameters to send to the API endpoint

	for the list projects operation.

	Typically these are written to a http.Request.
*/
type ListProjectsParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Name.

	   The name of project.
	*/
	Name *string

	/* Owner.

	   The name of project owner.
	*/
	Owner *string

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	/* Public.

	   The project is public or private.
	*/
	Public *bool

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string

	/* Sort.

	   Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with "sort=field1,-field2"
	*/
	Sort *string

	/* WithDetail.

	   Bool value indicating whether return detailed information of the project

	   Default: true
	*/
	WithDetail *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectsParams) WithDefaults() *ListProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)

		withDetailDefault = bool(true)
	)

	val := ListProjectsParams{
		Page:       &pageDefault,
		PageSize:   &pageSizeDefault,
		WithDetail: &withDetailDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list projects params
func (o *ListProjectsParams) WithTimeout(timeout time.Duration) *ListProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list projects params
func (o *ListProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list projects params
func (o *ListProjectsParams) WithContext(ctx context.Context) *ListProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list projects params
func (o *ListProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list projects params
func (o *ListProjectsParams) WithHTTPClient(client *http.Client) *ListProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list projects params
func (o *ListProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list projects params
func (o *ListProjectsParams) WithXRequestID(xRequestID *string) *ListProjectsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list projects params
func (o *ListProjectsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the list projects params
func (o *ListProjectsParams) WithName(name *string) *ListProjectsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list projects params
func (o *ListProjectsParams) SetName(name *string) {
	o.Name = name
}

// WithOwner adds the owner to the list projects params
func (o *ListProjectsParams) WithOwner(owner *string) *ListProjectsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list projects params
func (o *ListProjectsParams) SetOwner(owner *string) {
	o.Owner = owner
}

// WithPage adds the page to the list projects params
func (o *ListProjectsParams) WithPage(page *int64) *ListProjectsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list projects params
func (o *ListProjectsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list projects params
func (o *ListProjectsParams) WithPageSize(pageSize *int64) *ListProjectsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list projects params
func (o *ListProjectsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPublic adds the public to the list projects params
func (o *ListProjectsParams) WithPublic(public *bool) *ListProjectsParams {
	o.SetPublic(public)
	return o
}

// SetPublic adds the public to the list projects params
func (o *ListProjectsParams) SetPublic(public *bool) {
	o.Public = public
}

// WithQ adds the q to the list projects params
func (o *ListProjectsParams) WithQ(q *string) *ListProjectsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list projects params
func (o *ListProjectsParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the list projects params
func (o *ListProjectsParams) WithSort(sort *string) *ListProjectsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list projects params
func (o *ListProjectsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithWithDetail adds the withDetail to the list projects params
func (o *ListProjectsParams) WithWithDetail(withDetail *bool) *ListProjectsParams {
	o.SetWithDetail(withDetail)
	return o
}

// SetWithDetail adds the withDetail to the list projects params
func (o *ListProjectsParams) SetWithDetail(withDetail *bool) {
	o.WithDetail = withDetail
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Owner != nil {

		// query param owner
		var qrOwner string

		if o.Owner != nil {
			qrOwner = *o.Owner
		}
		qOwner := qrOwner
		if qOwner != "" {

			if err := r.SetQueryParam("owner", qOwner); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Public != nil {

		// query param public
		var qrPublic bool

		if o.Public != nil {
			qrPublic = *o.Public
		}
		qPublic := swag.FormatBool(qrPublic)
		if qPublic != "" {

			if err := r.SetQueryParam("public", qPublic); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.WithDetail != nil {

		// query param with_detail
		var qrWithDetail bool

		if o.WithDetail != nil {
			qrWithDetail = *o.WithDetail
		}
		qWithDetail := swag.FormatBool(qrWithDetail)
		if qWithDetail != "" {

			if err := r.SetQueryParam("with_detail", qWithDetail); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
