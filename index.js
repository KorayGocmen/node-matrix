"use strict";

const { exec } = require("child_process");

const TwoDArrayToString = (m) => {
  let mStr = `${m.length},${m[0].length},`;
  for(let i = 0; i < m.length; i++) {
    const row = m[i];
    for(let j = 0; j < row.length; j++){
      const val = row[j];
      mStr += `${val},`;
    }
  }
  return mStr.substring(0, mStr.length - 1);
}

const FindOsBinary = () => {
  switch (process.platform) {
    case "darwin": {
      return "node-matrix_darwin_amd64";
    }
    case "linux": {
      return "node-matrix_linux_amd64";
    }
    case "win32": {
      return "node-matrix_windows_amd64.exe";
    }
    case "freebsd": {
      return "node-matrix_freebsd_arm";
    }
    default: {
      return null;
    }
  }
}

const Dot = (matrix1, matrix2) => {
  return new Promise((resolve, reject) => {

    if(!matrix1.length || !matrix2.length){
      reject("empty matrix");
    } else if (matrix1[0].length !== matrix2.length){
      reject("can not dot multiply due to matrix dimensions");
    }

    const binary = FindOsBinary();
    if (!binary) {
      reject("unknown os library");
    }

    const command = `./${binary} -command=dot -input1=${TwoDArrayToString(matrix1)} -input2=${TwoDArrayToString(matrix2)}`

    exec(command, (err, stdout, stderr) => {
      if (err) {
        reject(err);
      } else {
        console.log(stdout);
        if(stderr){
          reject(stderr);
        } else {
          const resp = JSON.parse(stdout);
          if (resp.success) {
            resolve(resp.out.matrix)
          } else {
            reject(resp.err);
          }
        }
      }
    });
  });
};

module.exports = {
  Dot
};