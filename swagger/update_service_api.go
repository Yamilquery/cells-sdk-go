/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type UpdateServiceApiService service


/* UpdateServiceApiService Apply an update to a given version
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param targetVersion 
 @return UpdateApplyUpdateResponse*/
func (a *UpdateServiceApiService) ApplyUpdate(ctx context.Context, targetVersion string) (UpdateApplyUpdateResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  UpdateApplyUpdateResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/update/{TargetVersion}"
	localVarPath = strings.Replace(localVarPath, "{"+"TargetVersion"+"}", fmt.Sprintf("%v", targetVersion), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* UpdateServiceApiService Check the remote server to see if there are available binaries
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "channel" (string) Channel name.
     @param "packageName" (string) Name of the currently running application.
     @param "currentVersion" (string) Current version of the application.
     @param "gOOS" (string) Current GOOS.
     @param "gOARCH" (string) Current GOARCH.
     @param "serviceName" (string) Not Used : specific service to get updates for.
 @return UpdateUpdateResponse*/
func (a *UpdateServiceApiService) UpdateRequired(ctx context.Context, localVarOptionals map[string]interface{}) (UpdateUpdateResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  UpdateUpdateResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["channel"], "string", "channel"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["packageName"], "string", "packageName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["currentVersion"], "string", "currentVersion"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["gOOS"], "string", "gOOS"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["gOARCH"], "string", "gOARCH"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["serviceName"], "string", "serviceName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["channel"].(string); localVarOk {
		localVarQueryParams.Add("Channel", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["packageName"].(string); localVarOk {
		localVarQueryParams.Add("PackageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["currentVersion"].(string); localVarOk {
		localVarQueryParams.Add("CurrentVersion", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["gOOS"].(string); localVarOk {
		localVarQueryParams.Add("GOOS", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["gOARCH"].(string); localVarOk {
		localVarQueryParams.Add("GOARCH", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["serviceName"].(string); localVarOk {
		localVarQueryParams.Add("ServiceName", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

