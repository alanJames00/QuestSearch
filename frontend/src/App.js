import React, { useState } from "react";
import SearchBar from "./components/SearchBar";
import SearchResults from "./components/SearchResults";
import Dropdown from "./components/Dropdown";
import Pagination from "./components/Pagination";
import { searchQuestions } from "./api/grpc/grpcClient";

const App = () => {
	const [searchQuery, setSearchQuery] = useState("");

	const [searchMode, setSearchMode] = useState("normal");
	const [results, setResults] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [totalPages, setTotalPages] = useState(1);
	const [pageSize, setPageSize] = useState(10);
	const [filter, setFilter] = useState("ALL"); // question type filter

	// searchModeOptions
	const searchModeOptions = [
		{ value: "normal", label: "Normal DB Search" },
		{ value: "suggested", label: "Elastic Search" },
	];

	// question type filters
	const questionTypeOptions = [
		{ value: "ALL", label: "All Types" },
		{ value: "ANAGRAM", label: "Anagram" },
		{ value: "MCQ", label: "MCQ" },
		{ value: "CONVERSATION", label: "Conversation" },
		{ value: "READ_ALONG", label: "Read Along" },
		{ value: "CONTENT_ONLY", label: "Content Only" },
	];

	// get questions from grpcClient
	const fetchResults = async (query, qType, page) => {
		const results = await searchQuestions(query, qType, page, 10);

		console.log("grpcClient resp", results);
		setTotalPages(results.totalpages);
		setCurrentPage(page);
		setResults(results.questionsList);
	};

	// handle the search
	const handleSearch = (query) => {
		setSearchQuery(query);
		fetchResults(query, filter, 1);
	};

	// handle page change
	const handlePageChange = (page) => {
		setCurrentPage(page);
		fetchResults(searchQuery, filter, page); // Fetch again with the selected filter and page
	};

	return (
		<div className="min-h-screen bg-gray-100 p-8">
			<h1 className="text-2xl font-bold text-center mb-8">
				Search Bar with Suggestions
			</h1>

			<div className="flex justify-center mb-4 gap-4">
				<Dropdown
					options={searchModeOptions}
					selectedValue={searchMode}
					onChange={(e) => setSearchMode(e.target.value)}
				/>

				<Dropdown
					options={questionTypeOptions}
					selectedValue={filter}
					onChange={(e) => setFilter(e.target.value)}
				/>
			</div>

			<SearchBar
				suggestionsEnabled={searchMode === "suggested"}
				onSearch={handleSearch}
			/>

			<SearchResults results={results} />

			<div className="flex justify-center mt-8">
				<Pagination
					currentPage={currentPage}
					totalPages={totalPages}
					onPageChange={handlePageChange}
				/>
			</div>
		</div>
	);
};

export default App;
