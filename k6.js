import http from 'k6/http';
import { check, sleep } from 'k6';
import exec from 'k6/execution';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js';

export const options = {
  stages: [
    { duration: '10s', target: 200 },
    { duration: '10s', target: 1000 },
    { duration: '30s', target: 2000 },
    { duration: '10s', target: 1000 },
    { duration: '4s', target: 0 },
  ],
};

export default function () {
  const res = http.get('http://34.77.52.21/posts/api/post/getAll');
  // if(!check(res, { 'status must be 200': (r) => r.status == 200 })){
  //   exec.test.abort()
  // }
  sleep(1);
}


export function handleSummary(data) {
  console.log('Finished executing performance tests');
  
  return {
    'stdout': textSummary(data, { indent: ' ', enableColors: true }), // Show the text summary to stdout...
    'summary.txt': textSummary(data, { indent: ' '}), // and a JSON with all the details...
  };
}