const { QuestionServiceClient } = require("./proto/question_grpc_web_pb");
const {
	GetQuestionsRequest,
	SearchQuestionRequest,
} = require("./proto/question_pb");

var client = new QuestionServiceClient("http://localhost:8080");

const getAllQuestions = async (page, limit) => {
	return new Promise((resolve, reject) => {
		const request = new GetQuestionsRequest();
		request.setPage(page);
		request.setLimit(limit);

		client.getQuestions(request, {}, (err, resp) => {
			if (err) {
				reject("error fetching question", err.message);
			} else {
				resolve(resp.toObject());
			}
		});
	});
};

const searchQuestions = async (query, page, limit) => {
	return new Promise((resolve, reject) => {
		const request = new SearchQuestionRequest();
		request.setSearchTerm(query);
		request.setPage(page);
		request.setLimit(limit);

		client.searchQuestion(request, {}, (err, resp) => {
			if (err) {
				reject("error fetching question", err.message);
			} else {
				resolve(resp.toObject());
			}
		});
	});
};

module.exports = {
	searchQuestions,
	getAllQuestions,
};
