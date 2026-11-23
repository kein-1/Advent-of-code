//
//  main.swift
//  
//
//  Created by Kein Li on 11/19/26.
//

import Foundation
import Shared

func main() {
		let args = CommandLine.arguments

		guard args.count >= 2 else {
			print("usage: Year2019 <day>")
			exit(1)
		}

		guard let day = Int(args[1]) else {
			print("'\(args[1])' is not a valid day number")
			exit(1)
		}

		let extras = Array(args.dropFirst(2))
		let file = extras.first ?? "data.txt"
		guard let argStr = extras.dropFirst().first, let part = Part(rawValue: argStr) else {
			print("wrong or missing part provided")
			exit(1)
		}
		switch day {
		case 1:
			day1(filePath: file, part: part)
		case 2:
			day2(filePath: file, part: part)
		default:
			print("Day \(day) is not implemented")
			exit(1)
		}
}

main()
