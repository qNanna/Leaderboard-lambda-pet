package repository

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	config "slanj/common/config"
	types "slanj/common/steam/types"
)

type SteamRepository struct{}

func (s *SteamRepository) GetSteamPlayers(ids []string) []types.SteamPlayer {
	steamConfig := config.GetConfig().Params.Steam
	endpointString := steamConfig.Endpoints.GetUsers + strings.Join(ids[:], ",")
	endpoint := os.Expand(endpointString, func(prop string) string { return steamConfig.Key })

	response, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data := &types.SteamPlayersResponse{}
	json.NewDecoder(response.Body).Decode(data)

	return data.Response.Players
}
