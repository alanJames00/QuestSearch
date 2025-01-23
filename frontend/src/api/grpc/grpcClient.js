const { QuestionServiceClient } = require("./proto/question_grpc_web_pb");
const {
	GetQuestionsRequest,
	GetQuestionsResponse,
} = require("./proto/question_pb");

var client = new QuestionServiceClient("http://localhost:8080");

var request = new GetQuestionsRequest();
request.setPage(1);
request.setLimit(10);

const fetchQuestions = async (page, limit) => {
	return new Promise((resolve, reject) => {
		const request = new GetQuestionsRequest();
		request.setPage(page);
		request.setLimit(page);

		client.getQuestions(request, {}, (err, resp) => {
			if (err) {
				reject("error fetching question", err.message);
			} else {
				resolve(resp);
			}
		});
	});
};

module.exports = {
	fetchQuestions,
};
