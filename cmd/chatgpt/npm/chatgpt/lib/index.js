const spawn = require("child_process").spawn;
const os = require("os");
const path = require("path");

const platform = os.platform();
const arch = os.arch();

const ERR_NOT_SUPPORT = new Error("chatgpt does not support your platform");

const platformMap = {
  win32: {
    ia32: "chatgpt-windows-386",
    arm: "chatgpt-windows-arm",
    arm64: "chatgpt-windows-arm64",
    x64: "chatgpt-windows-amd64",
  },
  darwin: {
    arm64: "chatgpt-darwin-arm64",
    x64: "chatgpt-darwin-amd64",
  },
  linux: {
    ia32: "chatgpt-linux-386",
    arm: "chatgpt-linux-arm",
    arm64: "chatgpt-linux-arm64",
    x64: "chatgpt-linux-amd64",
    mips: "chatgpt-linux-mips",
    mipsel: "chatgpt-linux-mipsel",
    mips64: "chatgpt-linux-mips64",
    mips64el: "chatgpt-linux-mips64el",
  },
  freebsd: {
    ia32: "chatgpt-freebsd-386",
    arm: "chatgpt-freebsd-arm",
    arm64: "chatgpt-freebsd-arm64",
    x64: "chatgpt-freebsd-amd64",
  },
  openbsd: {
    ia32: "chatgpt-openbsd-386",
    arm: "chatgpt-openbsd-arm",
    arm64: "chatgpt-openbsd-arm64",
    x64: "chatgpt-openbsd-amd64",
  },
};

const archMap = platformMap[platform];

if (!archMap) throw ERR_NOT_SUPPORT;

const prebuildPackageName = archMap[arch];

if (!prebuildPackageName) throw ERR_NOT_SUPPORT;

const binaryPackageDir = path.dirname(
  require.resolve(`@axetroy/${prebuildPackageName}/package.json`)
);

const executableFileName = `chatgpt${platform === "win32" ? ".exe" : ""}`;

const executableFilePath = path.join(binaryPackageDir, executableFileName);

/**
 *
 * @param {Array<string>} argv
 * @param {SpawnOptionsWithoutStdio} [spawnOptions]
 * @returns
 */
function exec(argv, spawnOptions = {}) {
  const ps = spawn(executableFilePath, argv, {
    ...spawnOptions,
    stdout: "piped",
  });

  return ps;
}

module.exports.exec = exec;
