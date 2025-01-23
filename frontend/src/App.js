import logo from "./logo.svg";
import "./App.css";
import { useEffect, useState } from "react";
import { fetchQuestions } from "./api/grpc/grpcClient";

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
		<div className="App">
			<header className="App-header">
				<img src={logo} className="App-logo" alt="logo" />
				<p>
					Edit <code>src/App.js</code> and save to reload.
				</p>
				<a
					className="App-link"
					href="https://reactjs.org"
					target="_blank"
					rel="noopener noreferrer"
				>
					Learn React
				</a>
			</header>
		</div>
	);
}

export default App;
