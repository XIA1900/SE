import { request } from 'umi';

export async function getPost(params) {
  return request('/api/getPost', {
    params,
  });
}

export async function getCollection(params) {
  return request('/api/getCollection', {
    params,
  });
}

export async function getReply(params) {
  return request('/api/getReply', {
    params,
  });
}

export async function getLike(params) {
  return request('/api/getLike', {
    params,
  });
}

