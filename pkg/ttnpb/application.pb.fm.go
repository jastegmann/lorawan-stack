// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

var _ApplicationFieldPaths = [...]string{
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"ids.application_id",
	"name",
	"updated_at",
}

func (*Application) FieldMaskPaths() []string {
	ret := make([]string, len(_ApplicationFieldPaths))
	copy(ret, _ApplicationFieldPaths[:])
	return ret
}

func (dst *Application) SetFields(src *Application, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "attributes":
			dst.Attributes = src.Attributes
		case "contact_info":
			dst.ContactInfo = src.ContactInfo
		case "created_at":
			dst.CreatedAt = src.CreatedAt
		case "description":
			dst.Description = src.Description
		case "ids":
			dst.ApplicationIdentifiers = src.ApplicationIdentifiers
		case "ids.application_id":
			dst.ApplicationIdentifiers.SetFields(&src.ApplicationIdentifiers, _pathsWithoutPrefix("ids", paths)...)
		case "name":
			dst.Name = src.Name
		case "updated_at":
			dst.UpdatedAt = src.UpdatedAt
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _ApplicationsFieldPaths = [...]string{
	"applications",
}

func (*Applications) FieldMaskPaths() []string {
	ret := make([]string, len(_ApplicationsFieldPaths))
	copy(ret, _ApplicationsFieldPaths[:])
	return ret
}

func (dst *Applications) SetFields(src *Applications, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "applications":
			dst.Applications = src.Applications
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _GetApplicationRequestFieldPaths = [...]string{
	"application_ids",
	"application_ids.application_id",
	"field_mask",
}

func (*GetApplicationRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_GetApplicationRequestFieldPaths))
	copy(ret, _GetApplicationRequestFieldPaths[:])
	return ret
}

func (dst *GetApplicationRequest) SetFields(src *GetApplicationRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "application_ids":
			dst.ApplicationIdentifiers = src.ApplicationIdentifiers
		case "application_ids.application_id":
			dst.ApplicationIdentifiers.SetFields(&src.ApplicationIdentifiers, _pathsWithoutPrefix("application_ids", paths)...)
		case "field_mask":
			dst.FieldMask = src.FieldMask
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _ListApplicationsRequestFieldPaths = [...]string{
	"collaborator",
	"collaborator.organization_ids",
	"collaborator.organization_ids.organization_id",
	"collaborator.user_ids",
	"collaborator.user_ids.email",
	"collaborator.user_ids.user_id",
	"field_mask",
	"limit",
	"order",
	"page",
}

func (*ListApplicationsRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_ListApplicationsRequestFieldPaths))
	copy(ret, _ListApplicationsRequestFieldPaths[:])
	return ret
}

func (dst *ListApplicationsRequest) SetFields(src *ListApplicationsRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "collaborator":
			dst.Collaborator = src.Collaborator
		case "collaborator.organization_ids":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.organization_ids.organization_id":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids.email":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids.user_id":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "field_mask":
			dst.FieldMask = src.FieldMask
		case "limit":
			dst.Limit = src.Limit
		case "order":
			dst.Order = src.Order
		case "page":
			dst.Page = src.Page
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _CreateApplicationRequestFieldPaths = [...]string{
	"application",
	"application.attributes",
	"application.contact_info",
	"application.created_at",
	"application.description",
	"application.ids",
	"application.ids.application_id",
	"application.name",
	"application.updated_at",
	"collaborator",
	"collaborator.organization_ids",
	"collaborator.organization_ids.organization_id",
	"collaborator.user_ids",
	"collaborator.user_ids.email",
	"collaborator.user_ids.user_id",
}

func (*CreateApplicationRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_CreateApplicationRequestFieldPaths))
	copy(ret, _CreateApplicationRequestFieldPaths[:])
	return ret
}

func (dst *CreateApplicationRequest) SetFields(src *CreateApplicationRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "application":
			dst.Application = src.Application
		case "application.attributes":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.contact_info":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.created_at":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.description":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.ids":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.ids.application_id":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.name":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.updated_at":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "collaborator":
			dst.Collaborator = src.Collaborator
		case "collaborator.organization_ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.organization_ids.organization_id":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids.email":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids.user_id":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _UpdateApplicationRequestFieldPaths = [...]string{
	"application",
	"application.attributes",
	"application.contact_info",
	"application.created_at",
	"application.description",
	"application.ids",
	"application.ids.application_id",
	"application.name",
	"application.updated_at",
	"field_mask",
}

func (*UpdateApplicationRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_UpdateApplicationRequestFieldPaths))
	copy(ret, _UpdateApplicationRequestFieldPaths[:])
	return ret
}

func (dst *UpdateApplicationRequest) SetFields(src *UpdateApplicationRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "application":
			dst.Application = src.Application
		case "application.attributes":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.contact_info":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.created_at":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.description":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.ids":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.ids.application_id":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.name":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "application.updated_at":
			dst.Application.SetFields(&src.Application, _pathsWithoutPrefix("application", paths)...)
		case "field_mask":
			dst.FieldMask = src.FieldMask
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _CreateApplicationAPIKeyRequestFieldPaths = [...]string{
	"application_ids",
	"application_ids.application_id",
	"name",
	"rights",
}

func (*CreateApplicationAPIKeyRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_CreateApplicationAPIKeyRequestFieldPaths))
	copy(ret, _CreateApplicationAPIKeyRequestFieldPaths[:])
	return ret
}

func (dst *CreateApplicationAPIKeyRequest) SetFields(src *CreateApplicationAPIKeyRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "application_ids":
			dst.ApplicationIdentifiers = src.ApplicationIdentifiers
		case "application_ids.application_id":
			dst.ApplicationIdentifiers.SetFields(&src.ApplicationIdentifiers, _pathsWithoutPrefix("application_ids", paths)...)
		case "name":
			dst.Name = src.Name
		case "rights":
			dst.Rights = src.Rights
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _UpdateApplicationAPIKeyRequestFieldPaths = [...]string{
	"api_key",
	"api_key.id",
	"api_key.key",
	"api_key.name",
	"api_key.rights",
	"application_ids",
	"application_ids.application_id",
}

func (*UpdateApplicationAPIKeyRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_UpdateApplicationAPIKeyRequestFieldPaths))
	copy(ret, _UpdateApplicationAPIKeyRequestFieldPaths[:])
	return ret
}

func (dst *UpdateApplicationAPIKeyRequest) SetFields(src *UpdateApplicationAPIKeyRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "api_key":
			dst.APIKey = src.APIKey
		case "api_key.id":
			dst.APIKey.SetFields(&src.APIKey, _pathsWithoutPrefix("api_key", paths)...)
		case "api_key.key":
			dst.APIKey.SetFields(&src.APIKey, _pathsWithoutPrefix("api_key", paths)...)
		case "api_key.name":
			dst.APIKey.SetFields(&src.APIKey, _pathsWithoutPrefix("api_key", paths)...)
		case "api_key.rights":
			dst.APIKey.SetFields(&src.APIKey, _pathsWithoutPrefix("api_key", paths)...)
		case "application_ids":
			dst.ApplicationIdentifiers = src.ApplicationIdentifiers
		case "application_ids.application_id":
			dst.ApplicationIdentifiers.SetFields(&src.ApplicationIdentifiers, _pathsWithoutPrefix("application_ids", paths)...)
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _SetApplicationCollaboratorRequestFieldPaths = [...]string{
	"application_ids",
	"application_ids.application_id",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"collaborator.rights",
}

func (*SetApplicationCollaboratorRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_SetApplicationCollaboratorRequestFieldPaths))
	copy(ret, _SetApplicationCollaboratorRequestFieldPaths[:])
	return ret
}

func (dst *SetApplicationCollaboratorRequest) SetFields(src *SetApplicationCollaboratorRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "application_ids":
			dst.ApplicationIdentifiers = src.ApplicationIdentifiers
		case "application_ids.application_id":
			dst.ApplicationIdentifiers.SetFields(&src.ApplicationIdentifiers, _pathsWithoutPrefix("application_ids", paths)...)
		case "collaborator":
			dst.Collaborator = src.Collaborator
		case "collaborator.ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.organization_ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.organization_ids.organization_id":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.user_ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.user_ids.email":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.user_ids.user_id":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.rights":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}