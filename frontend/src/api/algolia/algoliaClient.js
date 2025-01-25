import { searchClient } from "@algolia/client-search";

const client = searchClient(
	process.env.REACT_APP_ALGOLIA_APP_ID,
	process.env.REACT_APP_ALGOLIA_API_KEY,
);

async function getSuggestions(query, qType) {
	return new Promise((resolve, reject) => {
		const searchParams = {
			query: query,
			page: 0,
			filter: qType != "ALL" ? `(type: ${qType})` : undefined,
		};
		client
			.searchSingleIndex({ indexName: "questions", searchParams })
			.then((response) => {
				// map the hits' title to suggestion
				const suggestions = response.hits.map((hit) =>
					hit.title.substring(0, 50),
				);
				resolve(suggestions);
			})
			.catch((error) => {
				console.error(error);
				reject(error);
			});
	});
}

export { getSuggestions };
