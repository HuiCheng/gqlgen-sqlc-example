package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"
	"idea2/graph/model"
	"idea2/pkgs/db"
)

// Auth is the resolver for the auth field.
func (r *deviceResolver) Auth(ctx context.Context, obj *db.Device) (*db.DeviceAuth, error) {
	panic(fmt.Errorf("not implemented: Auth - auth"))
}

// Type is the resolver for the type field.
func (r *deviceAuthResolver) Type(ctx context.Context, obj *db.DeviceAuth) (model.DeviceAuthType, error) {
	return model.DeviceAuthType(obj.Type), nil
}

// CreateDeviceAuth is the resolver for the createDeviceAuth field.
func (r *mutationResolver) CreateDeviceAuth(ctx context.Context, input *db.CreateDeviceAuthParams) (*db.DeviceAuth, error) {
	return r.dbQueries.CreateDeviceAuth(ctx, input)
}

// GetDevice is the resolver for the GetDevice field.
func (r *queryResolver) GetDevice(ctx context.Context, id int64) (*db.Device, error) {
	panic(fmt.Errorf("not implemented: GetDevice - GetDevice"))
}

// GetDevices is the resolver for the GetDevices field.
func (r *queryResolver) GetDevices(ctx context.Context) ([]db.Device, error) {
	panic(fmt.Errorf("not implemented: GetDevices - GetDevices"))
}

// GetDeviceAuth is the resolver for the GetDeviceAuth field.
func (r *queryResolver) GetDeviceAuth(ctx context.Context) (*db.DeviceAuth, error) {
	panic(fmt.Errorf("not implemented: GetDeviceAuth - GetDeviceAuth"))
}

// GetDeviceAuths is the resolver for the GetDeviceAuths field.
func (r *queryResolver) GetDeviceAuths(ctx context.Context) ([]db.DeviceAuth, error) {
	panic(fmt.Errorf("not implemented: GetDeviceAuths - GetDeviceAuths"))
}

// Type is the resolver for the type field.
func (r *createDeviceAuthParamsResolver) Type(ctx context.Context, obj *db.CreateDeviceAuthParams, data model.DeviceAuthType) error {
	obj.Type = string(data)
	return nil
}

// Device returns DeviceResolver implementation.
func (r *Resolver) Device() DeviceResolver { return &deviceResolver{r} }

// DeviceAuth returns DeviceAuthResolver implementation.
func (r *Resolver) DeviceAuth() DeviceAuthResolver { return &deviceAuthResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// CreateDeviceAuthParams returns CreateDeviceAuthParamsResolver implementation.
func (r *Resolver) CreateDeviceAuthParams() CreateDeviceAuthParamsResolver {
	return &createDeviceAuthParamsResolver{r}
}

type deviceResolver struct{ *Resolver }
type deviceAuthResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type createDeviceAuthParamsResolver struct{ *Resolver }
