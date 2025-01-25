import React, { useState, useEffect } from "react";
import { getSuggestions } from "../api/algolia/algoliaClient";

const SearchBar = ({ suggestionsEnabled = true, onSearch }) => {
	const [query, setQuery] = useState("");
	const [suggestions, setSuggestions] = useState([]);
	const [isFocused, setIsFocused] = useState(false);

	// Handle input change
	const handleInputChange = async (e) => {
		const value = e.target.value;
		setQuery(value);

		if (suggestionsEnabled && value.trim()) {
			const filteredSuggestions = await getSuggestions(value, "ALL");
			console.log("filteredSuggestions", filteredSuggestions);
			setSuggestions(filteredSuggestions);
		} else {
			setSuggestions([]);
		}
	};

	// Handle input focus
	const handleFocus = () => {
		setIsFocused(true);
	};

	// Handle input blur
	const handleBlur = () => {
		setTimeout(() => {
			setIsFocused(false);
		}, 200);
	};

	// Handle suggestion click
	const handleSuggestionClick = (suggestion) => {
		setQuery(suggestion);
		setSuggestions([]);
	};

	// Handle search button click
	const handleSearchClick = () => {
		if (onSearch) {
			onSearch(query); // Trigger the parent callback with the query
		}
	};

	return (
		<div className="max-w-md mx-auto mt-8 px-4 sm:px-0">
			<div className="relative flex flex-col sm:flex-row items-center space-y-2 sm:space-y-0 sm:space-x-2">
				<input
					type="text"
					placeholder="Search..."
					value={query}
					onChange={handleInputChange}
					onFocus={handleFocus}
					onBlur={handleBlur}
					className="w-full sm:w-auto flex-1 px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
				/>
				<button
					onClick={handleSearchClick}
					className="w-full sm:w-auto px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
				>
					{query.length > 0 ? "Search" : "Show All"}
				</button>
			</div>

			{suggestionsEnabled && isFocused && suggestions.length > 0 && (
				<div className="absolute left-0 right-0 bg-white border border-gray-300 rounded-lg mt-2 shadow-lg z-10 max-h-48 overflow-auto">
					{suggestions.slice(0, 10).map((suggestion, index) => (
						<div
							key={index}
							onClick={() => handleSuggestionClick(suggestion)}
							className="px-4 py-2 hover:bg-gray-100 cursor-pointer"
						>
							{suggestion}
						</div>
					))}
				</div>
			)}
		</div>
	);
};

export default SearchBar;
