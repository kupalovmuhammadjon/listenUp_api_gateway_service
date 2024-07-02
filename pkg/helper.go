package pkg

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
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func NewAuthenticationClient(cfg *config.Config) pbAuthentication.AuthenticationClient{
	conn, err := grpc.NewClient("localhost"+cfg.AUTHENTICATION_SERVICE_PORT, 
					grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting authentication service ", err)
	}
	a := pbAuthentication.NewAuthenticationClient(conn)

	return a
}

func NewCollaborationClient(cfg *config.Config) pbCollaboration.CollaborationsClient{
	conn, err := grpc.NewClient(cfg.COLLABORATIONS_SERVICE_PORT, 
				grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting collaborations service ", err)
	}
	a := pbCollaboration.NewCollaborationsClient(conn)

	return a
}

func NewCommentsClient(cfg *config.Config) pbComments.CommentsClient{
	conn, err := grpc.NewClient(cfg.COLLABORATIONS_SERVICE_PORT, 
				grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting collaborations service ", err)
	}
	a := pbComments.NewCommentsClient(conn)

	return a
}

func NewEpisodeMetadataClient(cfg *config.Config) pbEpisodeMetadata.EpisodeMetadataClient{
	conn, err := grpc.NewClient(cfg.DISCOVERY_SERVICE_PORT, 
				grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting discovery service ", err)
	}
	a := pbEpisodeMetadata.NewEpisodeMetadataClient(conn)

	return a
}

func NewUserInteractionsClient(cfg *config.Config) pbUserInteractions.UserInteractionsClient{
	conn, err := grpc.NewClient(cfg.DISCOVERY_SERVICE_PORT, 
				grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting discovery service ", err)
	}
	a := pbUserInteractions.NewUserInteractionsClient(conn)

	return a
}

func NewEpisodesClient(cfg *config.Config) pbEpisodes.EpisodesServiceClient{
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT, 
				grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting podcast service ", err)
	}
	a := pbEpisodes.NewEpisodesServiceClient(conn)

	return a
}


func NewPodcastsClient(cfg *config.Config) pbPodcasts.PodcastsClient{
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT, 
				grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting podcast service ", err)
	}
	a := pbPodcasts.NewPodcastsClient(conn)

	return a
}

func NewUserManagementClient(cfg *config.Config) pbUserManagement.UserManagementClient{
	conn, err := grpc.NewClient(cfg.USER_SERVICE_PORT, 
				grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting user service ", err)
	}
	a := pbUserManagement.NewUserManagementClient(conn)

	return a
}

