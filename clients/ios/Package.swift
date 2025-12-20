// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "PortfolioiOS",
    platforms: [
        .iOS(.v16)
    ],
    products: [
        .library(
            name: "PortfolioiOS",
            targets: ["PortfolioiOS"]
        ),
    ],
    targets: [
        .target(
            name: "PortfolioiOS",
            dependencies: []
        ),
        .testTarget(
            name: "PortfolioiOSTests",
            dependencies: ["PortfolioiOS"]
        ),
    ]
)

