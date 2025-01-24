import React from "react";

const Pagination = ({ currentPage, totalPages, onPageChange }) => {
	const handlePrevious = () => {
		if (currentPage > 1) {
			onPageChange(currentPage - 1);
		}
	};

	const handleNext = () => {
		if (currentPage < totalPages) {
			onPageChange(currentPage + 1);
		}
	};

	return (
		<div className="flex items-center gap-2">
			<button
				onClick={handlePrevious}
				disabled={currentPage === 1}
				className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:bg-gray-300"
			>
				Previous
			</button>
			<span className="mx-2">
				Page {currentPage} of {totalPages}
			</span>
			<button
				onClick={handleNext}
				disabled={currentPage === totalPages}
				className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:bg-gray-300"
			>
				Next
			</button>
		</div>
	);
};

export default Pagination;
