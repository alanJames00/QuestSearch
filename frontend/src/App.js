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
	const [loading, setLoading] = useState(false); // question loading state

	// searchModeOptions
	const searchModeOptions = [
		{ value: "normal", label: "Simple DB Search" },
		{ value: "suggested", label: "Algolia Search" },
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
		setLoading(true);

		try {
			const results = await searchQuestions(query, qType, page, 10);
			console.log("grpcClient resp", results);
			setTotalPages(results.totalpages);
			setCurrentPage(page);
			setResults(results.questionsList);
		} catch (err) {
			console.log("Error fetching results", err);
			alert("Error fetching results");
		} finally {
			setLoading(false);
		}
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
			<h1 className="text-2xl font-bold text-center mb-8">QuestSearch</h1>

			<div className="flex flex-wrap justify-center mb-4 gap-4">
				<div className="w-full sm:w-auto">
					<Dropdown
						options={searchModeOptions}
						selectedValue={searchMode}
						onChange={(e) => setSearchMode(e.target.value)}
					/>
				</div>
				<div className="w-full sm:w-auto">
					<Dropdown
						options={questionTypeOptions}
						selectedValue={filter}
						onChange={(e) => setFilter(e.target.value)}
					/>
				</div>
			</div>

			<SearchBar
				suggestionsEnabled={searchMode === "suggested"}
				onSearch={handleSearch}
			/>

			{loading ? (
				<div className="text-center mt-4">
					<span className="text-gray-600">Searching...</span>
				</div>
			) : (
				<SearchResults results={results} />
			)}

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
