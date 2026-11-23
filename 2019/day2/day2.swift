//
//  Day1.swift
//  
//
//  Created by Kein Li on 11/19/26.
//

import Foundation
import Shared


private enum Action: Int {
	case one = 1
	case two = 2
	case nineNine = 99
}

func day2(filePath: String, part: Part) {
	let content = try! readLines(filePath)
	guard content.count == 1 else { return }
	var input = content[0].split(separator: ",").map{ Int($0)!}
	switch part {
		case .one:
			part1(input)
		case .two:
			part2(input)
	}
}

@discardableResult
private func helper(_ input: [Int], _ val1: Int = 12, _ val2: Int = 0) -> [Int] {
	var input = input
	input[1] = val1
	input[2] = val2
	var ptr = 0

	while ptr <= input.count - 4 {
		let action = Action(rawValue: input[ptr])
		let inputLocationOne = input[ptr+1], inputLocationTwo = input[ptr+2], outputLocation = input[ptr+3]
		switch action {
			case .one:
				input[outputLocation] = input[inputLocationOne] + input[inputLocationTwo]
			case .two:
				input[outputLocation] = input[inputLocationOne] * input[inputLocationTwo]
			default:
				break
		}
		ptr+=4
	}
	return input
}


private func part1(_ input: [Int]) {
	helper(input)
}

private func part2(_ input: [Int]) {
	for i in 0..<100 {
		for j in 0..<100 {
			let newInput = helper(input,i,j)
//			print("____")
//			print(input,i,j)
//			print("____")
//			print(newInput)
//			if i == 0 && j >= 2 { return }

			if newInput[0] == 19690720 {
				print("We found it!\(newInput[1]) | \(newInput[j])! Solution is: \(100*newInput[1] + newInput[2])")
			}
		}
	}
}
