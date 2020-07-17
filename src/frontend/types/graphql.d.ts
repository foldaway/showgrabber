namespace GraphQLTypes {
  export interface Series {
    id: string;
    name: string;
    status: string;
    network: string;
    poster: string;
    tvdbID: number;
    overview: string;

    seasons: Season[];
  }

  export interface Season {
    id: string;
    number: number;
    episodes: Episode[];
  }

  export interface Episode {
    title: string;
    number: number;
    airDate: string;
  }

  export interface TVDBRating {
    average: number;
    count: number;
  }

  export interface TVDBImage {
    fileName: string;
    id: string;
    keyType: string;
    languageId: number;
    resolution: string;
    subKey: string;
    thumbnail: string;
    ratingsInfo: TVDBRating;
  }

  export interface TVDBSeries {
    added: string;
    addedBy: number;
    airsDayOfWeek: string;
    airsTime: string;
    aliases: string[];
    banner: string;
    firstAired: string;
    genre: string[];
    id: string;
    imdbId: string;
    lastUpdated: number;
    network: string;
    networkID: string;
    overview: string;
    rating: string;
    runtime: string;
    seriesId: string;
    seriesName: string;
    siteRating: number;
    siteRatingCount: number;
    status: string;
    zap2itId: string;

    summary: {
      airedEpisodes: string;
      airedSeasons: string[];
    };

    fanArtImages: TVDBImage[];
    posterImages: TVDBImage[];
    seasonImages: TVDBImage[];
    seasonWideImages: TVDBImage[];
    seriesImages: TVDBImage[];
  }
}
