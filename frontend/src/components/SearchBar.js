import React, { useState, useEffect } from "react";

const SearchBar = ({ suggestionsEnabled = true, onSearch }) => {
	const [query, setQuery] = useState("");
	const [suggestions, setSuggestions] = useState([]);
	const [isFocused, setIsFocused] = useState(false);

	const mockSuggestions = [
		"React",
		"JavaScript",
		"Tailwind CSS",
		"Node.js",
		"Python",
		"Django",
		"Flask",
		"Vue.js",
		"Angular",
		"Svelte",
	];

	const fetchSuggestions = (query) => {
		return new Promise((resolve) => {
			setTimeout(() => {
				resolve(
					mockSuggestions.filter((suggestion) =>
						suggestion.toLowerCase().includes(query.toLowerCase()),
					),
				);
			}, 300);
		});
	};

	// Handle input change
	const handleInputChange = async (e) => {
		const value = e.target.value;
		setQuery(value);

		if (suggestionsEnabled && value.trim()) {
			const filteredSuggestions = await fetchSuggestions(value);
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
		<div className="max-w-md mx-auto mt-8">
			<div className="relative flex items-center space-x-2">
				<input
					type="text"
					placeholder="Search..."
					value={query}
					onChange={handleInputChange}
					onFocus={handleFocus}
					onBlur={handleBlur}
					className="w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
				/>
				<button
					onClick={handleSearchClick}
					className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
				>
					Search
				</button>
			</div>

			{suggestionsEnabled && isFocused && suggestions.length > 0 && (
				<div className="absolute bg-white border border-gray-300 rounded-lg mt-1 shadow-lg z-10">
					{suggestions.map((suggestion, index) => (
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
