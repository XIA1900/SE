import { request } from 'umi';

export async function getPost(params) {
  return request('/api/getPost', {
    params,
  });
}
