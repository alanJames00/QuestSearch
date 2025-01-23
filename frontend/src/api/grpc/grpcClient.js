const { QuestionServiceClient } = require("./proto/question_grpc_web_pb");
const {
	GetQuestionsRequest,
	GetQuestionsResponse,
} = require("./proto/question_pb");

var client = new QuestionServiceClient("http://localhost:8080");

const fetchQuestions = async (page, limit) => {
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

module.exports = {
	fetchQuestions,
};
