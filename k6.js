import http from 'k6/http';
import { check, sleep } from 'k6';
import exec from 'k6/execution';

export const options = {
  stages: [
    { duration: '3s', target: 200 },
    { duration: '2s', target: 100 },
    { duration: '2s', target: 0 },
  ],
};

export default function () {
  const res = http.get('http://34.77.52.21/posts/api/post/');
  if(!check(res, { 'status must be 200': (r) => r.status == 200 })){
    exec.test.abort()
  }
  sleep(1);
}


export function handleSummary(data) {
  console.log('Finished executing performance tests');

  return {
    'stdout': textSummary(data, { indent: ' ', enableColors: true }), // Show the text summary to stdout...
  };
}