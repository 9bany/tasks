package db

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/9bany/task/util"
	"github.com/stretchr/testify/require"
)

func compareBytesOfMap(t *testing.T, byte1, byte2 []byte) {
	var map1 map[string]interface{}
	var map2 map[string]interface{}

	err := json.Unmarshal(byte1, &map1)
	require.NoError(t, err)

	err = json.Unmarshal(byte2, &map2)
	require.NoError(t, err)

	require.Equal(t, true, reflect.DeepEqual(map1, map2))

}

func createRandomSite(t *testing.T) Sites {

	jsonMap := util.RanJsonMap()
	data, err := json.Marshal(jsonMap)

	require.NoError(t, err)

	arg := CreateSiteParams{
		Url:      util.RandomString(20),
		MetaData: data,
	}

	site, err := testQueries.CreateSite(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, site)

	require.Equal(t, arg.Url, site.Url)

	compareBytesOfMap(t, arg.MetaData, site.MetaData)

	require.NotZero(t, site.CreatedAt)

	return site
}

func TestCreateSite(t *testing.T) {
	createRandomSite(t)
}

func TestGetSite(t *testing.T) {
	site0 := createRandomSite(t)

	site, err := testQueries.GetSiteByURL(context.Background(), site0.Url)

	require.NoError(t, err)
	require.NotEmpty(t, site)
	
	require.Equal(t, site0.Url, site.Url)
	require.Equal(t, site0.ID, site.ID)
	compareBytesOfMap(t, site0.MetaData, site.MetaData)
}
