// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organization pipelines API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization pipelines API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreatePipeline Create a new pipeline
*/
func (a *Client) CreatePipeline(params *CreatePipelineParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePipelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPipeline",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePipelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreatePipelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeletePipeline Delete the pipeline.
*/
func (a *Client) DeletePipeline(params *DeletePipelineParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePipelineNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePipeline",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePipelineNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePipelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DiffPipeline The diff between the provided pipeline configuration and the pipeline from the given name.
*/
func (a *Client) DiffPipeline(params *DiffPipelineParams, authInfo runtime.ClientAuthInfoWriter) (*DiffPipelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDiffPipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "diffPipeline",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DiffPipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DiffPipelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DiffPipelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPipeline Get the configuration of the pipeline.
*/
func (a *Client) GetPipeline(params *GetPipelineParams, authInfo runtime.ClientAuthInfoWriter) (*GetPipelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPipeline",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPipelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPipelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPipelineConfig Get the YAML configuration of the pipeline.
*/
func (a *Client) GetPipelineConfig(params *GetPipelineConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetPipelineConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPipelineConfig",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPipelineConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPipelineConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPipelineConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPipelineVariables Get the YAML variables of the pipeline.
*/
func (a *Client) GetPipelineVariables(params *GetPipelineVariablesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPipelineVariablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineVariablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPipelineVariables",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPipelineVariablesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPipelineVariablesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPipelineVariablesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPipelines Get all the pipelines that the authenticated user has access to.
*/
func (a *Client) GetPipelines(params *GetPipelinesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPipelinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPipelines",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/pipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPipelinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPipelinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPipelinesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProjectPipelines Get the pipelines that the authenticated user has access to.
*/
func (a *Client) GetProjectPipelines(params *GetProjectPipelinesParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectPipelinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectPipelinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProjectPipelines",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectPipelinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectPipelinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProjectPipelinesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PausePipeline Pause a pipeline
*/
func (a *Client) PausePipeline(params *PausePipelineParams, authInfo runtime.ClientAuthInfoWriter) (*PausePipelineNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPausePipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pausePipeline",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PausePipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PausePipelineNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PausePipelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RenamePipeline Rename a pipeline
*/
func (a *Client) RenamePipeline(params *RenamePipelineParams, authInfo runtime.ClientAuthInfoWriter) (*RenamePipelineNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRenamePipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "renamePipeline",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/rename",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RenamePipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RenamePipelineNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RenamePipelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UnpausePipeline Unpause a pipeline
*/
func (a *Client) UnpausePipeline(params *UnpausePipelineParams, authInfo runtime.ClientAuthInfoWriter) (*UnpausePipelineNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnpausePipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unpausePipeline",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/unpause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnpausePipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnpausePipelineNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnpausePipelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdatePipeline Update the configuration of the given pipeline name.
*/
func (a *Client) UpdatePipeline(params *UpdatePipelineParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePipelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePipeline",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePipelineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePipelineDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}