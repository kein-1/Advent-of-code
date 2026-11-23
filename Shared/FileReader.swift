//
//  File.swift
//  advent_of_code
//
//  Created by Kein Li on 11/19/26.
//

import Foundation

public func readLines(_ fileName: String, from sourceFile: String = #filePath) throws -> [String] {
		let folder = URL(fileURLWithPath: sourceFile).deletingLastPathComponent()
		let url = folder.appendingPathComponent(fileName)
		let text = try String(contentsOf: url, encoding: .utf8)
		var lines = text.split(omittingEmptySubsequences: false, whereSeparator: \.isNewline).map(String.init)
		if lines.last?.isEmpty == true { lines.removeLast() }
		return lines
}


public enum Part: String {
	case one
	case two
}
