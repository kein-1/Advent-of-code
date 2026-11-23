//
//  Day1.swift
//  
//
//  Created by Kein Li on 11/19/26.
//

import Foundation
import Shared


func day1(filePath: String, part: Part) {
	let content = try! readLines(filePath)
	switch part {
		case .one:
			part1(content)
		case .two:
			part2(content)
	}
}

private func part1(_ content: [String]) {
	var sum = 0
	for line in content {
		guard let numb = Int(line) else { continue }
		sum+=calcMass(numb, part: .one)
	}

	print(sum)
}

private func part2(_ content: [String]) {
	var sum = 0
	for line in content {
		guard let numb = Int(line) else { continue }
		let result = calcMass(numb, part: .two)
		sum+=result
		print("The fuel required by mass \(numb): \(result)")
	}
	print(sum)
}


private func calcMass(_ mass: Int, part: Part) -> Int {
	if mass <= 0 {
		return 0
	}
	let remainingFuel = max(0,Int((Double(mass)/3).rounded(.down) - 2))
	if part == .one {
		return remainingFuel
	}
	return remainingFuel + calcMass(remainingFuel,part: .two)
}
