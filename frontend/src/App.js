import logo from "./logo.svg";
import "./App.css";
import { useEffect, useState } from "react";
import { fetchQuestions } from "./api/grpc/grpcClient";
import SearchBar from "./components/searchBar";
import Dropdown from "./components/Dropdown";

function App() {
	const [questions, setQuestions] = useState([]);
	console.log(questions);

	async function getQuestions(page, limit) {
		try {
			const resp = await fetchQuestions(page, limit);
			console.log("questions fetched", resp);
			setQuestions(resp);
		} catch (err) {
			console.log("failed to fetch questions", err);
		}
	}

	useEffect(() => {
		getQuestions(1, 10);
	}, []);

	return (
		<div className="min-h-screen bg-gray-100 p-8">
			<h1 className="text-2xl font-bold text-center mb-8">
				QuestSearch: search any questions
			</h1>
			<SearchBar suggestionsEnabled={false} />{" "}
			{/* Set `false` to disable suggestions */}
		</div>
	);
}

export default App;
