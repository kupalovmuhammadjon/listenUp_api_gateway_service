package handler

import (
	"api_gateway/config"
	pbAuthentication "api_gateway/genproto/authentication"
	pbCollaboration "api_gateway/genproto/collaborations"
	pbComments "api_gateway/genproto/comments"
	pbEpisodeMetadata "api_gateway/genproto/episode_metadata"
	pbEpisodes "api_gateway/genproto/episodes"
	pbPodcasts "api_gateway/genproto/podcasts"
	pbUserManagement "api_gateway/genproto/user"
	pbUserInteractions "api_gateway/genproto/user_interactions"
	"api_gateway/pkg"
)

type Handler struct {
	ClientAuthentication   pbAuthentication.AuthenticationClient
	ClientCollaboration    pbCollaboration.CollaborationsClient
	ClientComments         pbComments.CommentsClient
	ClientEpisodeMetadata  pbEpisodeMetadata.EpisodeMetadataClient
	ClientEpisodes         pbEpisodes.EpisodesServiceClient
	ClientPodcasts         pbPodcasts.PodcastsClient
	ClientUserManagement   pbUserManagement.UserManagementClient
	ClientUserInteractions pbUserInteractions.UserInteractionsClient
}

func NewHandler(cfg *config.Config) *Handler {
	
	return &Handler{
		ClientAuthentication: pkg.NewAuthenticationClient(cfg), 
		ClientCollaboration:   pkg.NewCollaborationClient(cfg),
		ClientComments: pkg.NewCommentsClient(cfg),
		ClientEpisodeMetadata: pkg.NewEpisodeMetadataClient(cfg),
		ClientEpisodes: pkg.NewEpisodesClient(cfg),
		ClientPodcasts: pkg.NewPodcastsClient(cfg),
		ClientUserManagement: pkg.NewUserManagementClient(cfg),
		ClientUserInteractions: pkg.NewUserInteractionsClient(cfg),
	}
}
