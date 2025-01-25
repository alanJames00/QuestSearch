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
							{result.type === "ANAGRAM" && (
								<div className="mt-2">
									<ul className="mt-2 space-y-1">
										{result.anagram.blocksList.map((option, index) => (
											<li key={index} className="text-gray-700">
												<span className="font-bold">
													{String.fromCharCode(65 + index)}.
												</span>{" "}
												{option.text}
											</li>
										))}
									</ul>
								</div>
							)}
							{result.type === "MCQ" && (
								<div className="mt-2">
									<ul className="mt-2 space-y-1">
										{result.mcq.optionsList.map((option, index) => (
											<li key={index} className="text-gray-700">
												<span className="font-bold">
													{String.fromCharCode(65 + index)}.
												</span>{" "}
												{option.text}
											</li>
										))}
									</ul>
								</div>
							)}
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
