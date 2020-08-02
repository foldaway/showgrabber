namespace GraphQLTypes {
  export interface Series {
    id: string;
    name: string;
    status: string;
    network: string;
    poster: string;
    tvdbID: number;
    overview: string;
    language: Language;

    seasons: Season[];
  }

  export interface Season {
    id: string;
    number: number;
    episodes: Episode[];
  }

  export interface Episode {
    id: string;
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

  export interface NewznabComment {
    title: string;
    content: string;
    pub_date: Date;
  }

  export interface Newznab {
    id: string;
    title: string;
    description: string;
    size: number;
    air_date: Date;
    pub_date: Date;
    usenet_date: Time;
    num_grabs: number;
    num_comments: number;
    comments: NewznabComment[];

    source_endpoint: string;
    source_apikey: string;
    category: string[];
    info: string;
    genre: string;

    resolution: string;

    tvdbid: string;
    tvrageid: string;
    tvmazeid: string;
    season: string;
    episode: string;
    tvtitle: string;
    rating: number;

    imdb: string;
    imdbtitle: string;
    imdbyear: number;
    imdbscore: number;
    coverurl: string;

    seeders: number;
    peers: number;
    infohash: string;
    download_url: string;
    is_torrent: boolean;
  }

  export interface Language {
    abbreviation: string;
    englishName: string;
    tvdbID: number;
    name: string;
  }
}
