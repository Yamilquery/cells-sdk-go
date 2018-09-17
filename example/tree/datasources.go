package tree

import (
	"github.com/pydio/cells-sdk-go/client/config_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

// ListDatasources returns a collection of known datasources using the current default connection.
func ListDatasources() (*models.RestDataSourceCollection, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return nil, err
	}

	params := config_service.NewListDataSourcesParamsWithContext(ctx)
	result, err := apiClient.ConfigService.ListDataSources(params)
	if err != nil {
		return nil, err
	}

	return result.Payload, nil
}

// AddLocalDatasource creates a new local file based datasource.
func AddLocalDatasource(name, peerAddress string, port int32, rootFolderAbsPath string) (*models.ObjectDataSource, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return nil, err
	}

	ods := &models.ObjectDataSource{}
	ods.StorageType = "LOCAL"
	ods.Name = name
	ods.PeerAddress = peerAddress
	ods.ObjectsPort = port
	ods.StorageConfiguration = map[string]string{
		"folder":    rootFolderAbsPath,
		"normalize": "false",
	}

	params := config_service.NewPutDataSourceParamsWithContext(ctx)
	params.Name = name
	params.Body = ods

	result, err := apiClient.ConfigService.PutDataSource(params)
	if err != nil {
		return nil, err
	}

	return result.Payload, nil
}

// GetDatasource returns a datasource using the current default connection.
func GetDatasource(dsName string) (*models.ObjectDataSource, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return nil, err
	}

	params := config_service.NewGetDataSourceParamsWithContext(ctx)
	params.Name = dsName
	result, err := apiClient.ConfigService.GetDataSource(params)
	if err != nil {
		return nil, err
	}

	return result.Payload, nil
}

// DeleteLocalDatasource removes a local file based datasource, without removing underlying files.
func DeleteLocalDatasource(name string) (bool, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return false, err
	}

	params := config_service.NewDeleteDataSourceParamsWithContext(ctx)
	params.Name = name

	result, err := apiClient.ConfigService.DeleteDataSource(params)
	if err != nil {
		return false, err
	}

	return result.Payload.Success, nil
}
