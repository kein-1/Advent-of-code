// swift-tools-version: 6.2
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription


let package = Package(
		name: "advent_of_code",
		platforms: [.macOS(.v13)],
		targets: [
			.target(name: "Shared", path: "Shared"),
			.executableTarget(name: "Year2019", dependencies: ["Shared"], path: "2019")
		]
)
