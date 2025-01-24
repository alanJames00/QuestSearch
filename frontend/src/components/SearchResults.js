import React from "react";

const SearchResults = ({ results }) => {
	return (
		<div className="mt-8">
			{results.length > 0 ? (
				<ul className="space-y-4">
					{results.map((result) => (
						<li key={result.id} className="p-4 bg-white rounded-lg shadow">
							<h3 className="text-xl font-semibold">{result.title}</h3>
							<p className="text-gray-600">Type: {result.type}</p>
						</li>
					))}
				</ul>
			) : (
				<p className="text-center text-gray-600">No results found.</p>
			)}
		</div>
	);
};

export default SearchResults;
