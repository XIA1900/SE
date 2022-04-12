export async function uploadLogoImg(body, options) {
    console.log(body);
    return request('/api/user/login', {
      method: 'POST',
      data: body,
    });
  }