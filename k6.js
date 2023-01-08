import http from 'k6/http';
import { check, sleep } from 'k6';
import exec from 'k6/execution';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js';

export const options = {
  stages: [
    { duration: '500s', target: 5000 },
  ],
};

export default function () {
  const res = http.get('http://34.77.52.21/posts/api/post/getAll');
  // if(!check(res, { 'status must be 20000': (r) => r.status == 200 })){
  //   exec.test.abort()
  // }
  sleep(5);
}


export function handleSummary(data) {
  console.log('Finished executing performance tests');
  
  return {
    'stdout': textSummary(data, { indent: ' ', enableColors: true }), // Show the text summary to stdout...
    'summary.txt': textSummary(data, { indent: ' '}), // and a JSON with all the details...
  };
}