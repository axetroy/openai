/**
 * Usage:
 *
 * GIT_REF=refs/tags/v1.0.0 node npm/prepare.js
 */
const fs = require("fs");
const path = require("path");

const ref = process.env.GIT_REF; // refs/tags/v1.0.0

const arr = ref.split("/");
const version = arr[arr.length - 1].replace(/^v/, "");

console.log(`prepare publish to npm for: ${version}`);

const packages = fs
  .readdirSync(__dirname)
  .filter((v) => v.startsWith("chatgpt-"))
  .concat(["chatgpt"]);

for (const pkgName of packages) {
  const pkgPath = path.join(__dirname, pkgName, "package.json");

  const pkg = require(pkgPath);

  pkg.version = version;

  if (pkg.optionalDependencies) {
    for (const subDeps in pkg.optionalDependencies) {
      if (subDeps.startsWith("@axetroy/chatgpt-")) {
        pkg.optionalDependencies[subDeps] = version;
      }
    }
  }

  fs.writeFileSync(pkgPath, JSON.stringify(pkg, null, 2));

  if (pkgName.startsWith("chatgpt-")) {
    const fileMap = {
      "chatgpt-darwin-arm64": "chatgpt_darwin_arm64",
      "chatgpt-darwin-amd64": "chatgpt_darwin_amd64_v1",
      "chatgpt-linux-386": "chatgpt_linux_386",
      "chatgpt-linux-arm": "chatgpt_linux_arm_7",
      "chatgpt-linux-amd64": "chatgpt_linux_amd64_v1",
      "chatgpt-linux-arm64": "chatgpt_linux_arm64",
      "chatgpt-linux-mips": "chatgpt_linux_mips_softfloat",
      "chatgpt-linux-mipsel": "chatgpt_linux_mipsle_softfloat",
      "chatgpt-linux-mips64": "chatgpt_linux_mips64_softfloat",
      "chatgpt-linux-mips64el": "chatgpt_linux_mips64le_softfloat",
      "chatgpt-freebsd-386": "chatgpt_freebsd_386",
      "chatgpt-freebsd-arm": "chatgpt_freebsd_arm_7",
      "chatgpt-freebsd-arm64": "chatgpt_freebsd_arm64",
      "chatgpt-freebsd-amd64": "chatgpt_freebsd_amd64_v1",
      "chatgpt-openbsd-386": "chatgpt_openbsd_386",
      "chatgpt-openbsd-arm": "chatgpt_openbsd_arm_7",
      "chatgpt-openbsd-arm64": "chatgpt_openbsd_arm64",
      "chatgpt-openbsd-amd64": "chatgpt_openbsd_amd64_v1",
      "chatgpt-windows-386": "chatgpt_windows_386",
      "chatgpt-windows-amd64": "chatgpt_windows_amd64_v1",
      "chatgpt-windows-arm": "chatgpt_windows_arm_7",
      "chatgpt-windows-arm64": "chatgpt_windows_arm64",
    };

    if (pkgName in fileMap === false)
      throw new Error(`Can not found prebuild file for package '${pkgName}'`);

    const distFolder = fileMap[pkgName];

    const executableFileName =
      "chatgpt" + (pkgName.indexOf("windows") > -1 ? ".exe" : "");

    const executableFilePath = path.join(
      __dirname,
      "..",
      "dist",
      distFolder,
      executableFileName
    );

    fs.copyFileSync(
      executableFilePath,
      path.join(__dirname, pkgName, executableFileName)
    );
  } else {
    fs.copyFileSync(
      path.join(__dirname, "..", "README.md"),
      path.join(__dirname, pkgName, "README.md")
    );

    fs.copyFileSync(
      path.join(__dirname, "..", "LICENSE"),
      path.join(__dirname, pkgName, "LICENSE")
    );
  }
}
