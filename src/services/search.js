import { request } from 'umi';

export async function search(values) {
  return request('/api/queryList', {
    params,
  });
}
