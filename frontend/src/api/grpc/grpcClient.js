const { QuestionServiceClient } = require("./proto/question_grpc_web_pb");
const {
	GetQuestionsRequest,
	SearchQuestionRequest,
} = require("./proto/question_pb");

var client = new QuestionServiceClient(process.env.REACT_APP_GRPC_PROXY_HOST);

const getAllQuestions = async (qType, page, limit) => {
	return new Promise((resolve, reject) => {
		const request = new GetQuestionsRequest();
		request.setQType(qType);
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

const searchQuestions = async (query, qType, page, limit) => {
	return new Promise((resolve, reject) => {
		const request = new SearchQuestionRequest();
		request.setSearchTerm(query);
		request.setQType(qType);
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
