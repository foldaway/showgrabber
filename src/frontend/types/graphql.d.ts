namespace GraphQLTypes {
  export interface Series {
    id: string;
    name: string;
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

    fanArtImages: TVDBImage[];
    posterImages: TVDBImage[];
    seasonImages: TVDBImage[];
    seasonWideImages: TVDBImage[];
    seriesImages: TVDBImage[];
  }
}