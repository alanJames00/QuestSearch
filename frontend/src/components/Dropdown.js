import React from "react";

const Dropdown = ({ options, selectedValue, onChange, className }) => {
	return (
		<select
			value={selectedValue}
			onChange={onChange}
			className={`px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 ${className}`}
		>
			{options.map((option) => (
				<option key={option.value} value={option.value}>
					{option.label}
				</option>
			))}
		</select>
	);
};

export default Dropdown;
