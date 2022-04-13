import React from 'react';
import { PropertySafetyOutlined, UploadOutlined } from '@ant-design/icons';
import { Button, Input, Upload, message } from 'antd';
import ProForm, {
  ProFormDependency,
  ProFormFieldSet,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-form';
import { useRequest, history, useIntl} from 'umi';
import { queryCurrent, userUpdate } from '@/services/user';
import { uploadLogoImg } from '@/services/upload';
import styles from './BaseView.less';

const username = history.location.search.substring(1);

const props = {
  name: "uploadFilename",
  action: '/api/file/upload',
  headers: {
    authorization: 'authorization-text',
  },
  onChange(info) {
    if (info.file.status !== 'uploading') {
      console.log(info.file, info.fileList);
    }
    if (info.file.status === 'done') {
      message.success(`${info.file.name} file uploaded successfully`);
      history.push({
        pathname: '/account/settings',
        search:username,
      });
    } else if (info.file.status === 'error') {
      message.error(`${info.file.name} file upload failed.`);
    }
  },
  //设置只上传一张图片，根据实际情况修改
  // customRequest: async info => {
  //   // 手动上传
  //   const formData = new FormData();
  //   formData.append('uploadFilename', info.file);
  //   const result = await uploadLogoImg(formData);
  //   console.log(result);
  //   if(result.code === 200) {
  //     console.log("upload successful");
  //     history.push({
  //       pathname: '/account/settings',
  //       search:username,
  //     });
  //   }
  // },
  onRemove: file => {
    // 删除图片调用
    this.setState(state => {
      const index = state.fileList.indexOf(file);
      const newFileList = state.fileList.slice();
      newFileList.splice(index, 1);
      return {
        fileList: newFileList,
      };
    });
  },

  beforeUpload: file => {
    // 控制上传图片格式
    const isJpgOrPng = file.type === 'image/png';

    if (!isJpgOrPng) {
      message.error('Only JPEG/PNG files are allowed!');
      return;
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error('File size must be smaller than 2MB!');
      return;
    }
  },
}

const validatorPhone = (rule, value, callback) => {
  if (!value[0]) {
    callback('Please input your area code!');
  }

  if (!value[1]) {
    callback('Please input your phone number!');
  }

  callback();
}; 

const AvatarView = ({ avatar }) => (
  <>
    <div className={styles.avatar_title}></div>
    <div className={styles.avatar}>
      <img src={avatar} alt="avatar" />
    </div>
    <Upload {...props}>
      <div className={styles.button_view}>
        <Button>
          <UploadOutlined />
          change your avatar
        </Button>
      </div>
    </Upload>
  </>
);

const BaseView = () => {
  const intl = useIntl();

  const { data, loading } = useRequest(
    async() => {
      const result = await queryCurrent({
        username: username,
      });
      return result;
    },
    {
      formatResult: result => result,
    }
  );

  console.log(data);
  let currentUser = [];
  if(typeof(data) != 'undefined') {
    currentUser = {
      name: data.Username,
      birthday: data.Birthday,
      gender: data.Gender,
      major: data.Department,
      avatar: 'http://10.20.0.166:10010/resources/userfiles/'+ data.Username+'/avatar.png',
    };
  }

  const getAvatarURL = () => {
    if (currentUser) {
      if (currentUser.avatar) {
        return currentUser.avatar;
      }

      const url = 'https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png';
      return url;
    }

    return '';
  };

  const handleFinish = async (values) => {
    const data = {
      Username: values.username,
      Birthday: values.birthday,
      Gender: values.gender,
      Department: values.major,
    }
    const result = await userUpdate(data);
    if(result.code === 200)  {
      const defaultupdateInfoMessage = intl.formatMessage({
        id: 'updateInfo',
        defaultMessage: 'Update Successfully',
      });
      message.success(defaultupdateInfoMessage);
    }
  };

  return (
    <div className={styles.baseView}>
      {loading ? null : (
        <>
          <div className={styles.left}>
            <ProForm
              layout="vertical"
              onFinish={handleFinish}
              submitter={{
                resetButtonProps: {
                  style: {
                    display: 'none',
                  },
                },
                submitButtonProps: {
                  children: '更新基本信息',
                },
              }}
              initialValues={{}}
              hideRequiredMark
              
            >
              <ProForm.Group>
                <ProFormText
                  width="md"
                  name="username"
                  label="Username"
                  rules={[
                    {
                      required: true,
                      message: 'Please input your username!',
                    },
                  ]}
                  initialValue={currentUser.name}
                  layout="inline"
                  disabled
                />
                <ProFormText
                  width="md"
                  name="email"
                  label="Email"
                  rules={[
                    {
                      required: true,
                      message: 'Please input your email address!',
                    },
                  ]}
                  initialValue={currentUser.email}
                  layout="inline"
                />
                
              </ProForm.Group>
              
              <ProFormSelect
                width="md"
                name="gender"
                label="Gender"
                rules={[
                  {
                    required: true,
                    message: 'Please input your country!',
                  },
                ]}
                options={[
                  {
                    label: 'female',
                    value: 'Female',
                  },
                  {
                    label: 'male',
                    value: 'Male',
                  },
                  {
                    label: 'hide',
                    value: 'Prefer not to say',
                  }
                ]}
                initialValue = {currentUser.sex}
                layout="inline"
              />

              <ProFormText
                width="md"
                name="birthday"
                label="Birthday"
                rules={[
                  {
                    required: true,
                    message: 'Please input your birthday!',
                  },
                ]}
                initialValue={currentUser.birthday}
              />
              <ProFormText
                width="md"
                name="major"
                label="Major"
                rules={[
                  {
                    required: false,
                    message: 'Please input your major!',
                  },
                ]}
                initialValue={currentUser.major}
                display="inline-block"
              />
              <ProFormText
                width="md"
                name="grade"
                label="Grade"
                rules={[
                  {
                    required: false,
                    message: 'Please input your grade!',
                  },
                ]}
                initialValue={currentUser.grade}
                display="inline-block"
              />
              <ProFormTextArea
                name="signature"
                label="Signature"
                rules={[
                  {
                    required: false,
                    message: 'Please input your profile!',
                  },
                ]}
                placeholder="Tomorrow is another day."
                initialValue={currentUser.signature}
              />
              <ProFormSelect
                width="sm"
                name="country"
                label="Country"
                rules={[
                  {
                    required: true,
                    message: 'Please input your country!',
                  },
                ]}
                options={[
                  {
                    label: 'United States',
                    value: 'United States',
                  },
                  {
                    label: 'China',
                    value: 'China',
                  },
                ]}
                initialValue={currentUser.country}
              />

              <ProFormSelect
                width="sm"
                name="state"
                label="State"
                rules={[
                  {
                    required: true,
                    message: 'Please input your state!',
                  },
                ]}
                options={[
                  {
                    label: 'Florida',
                    value: 'Florida',
                  },
                  {
                    label: 'Texas',
                    value: 'Texas',
                  },
                ]}
                initialValue={currentUser.province}
              />

              <ProFormSelect
                width="sm"
                name="city"
                label="City"
                rules={[
                  {
                    required: true,
                    message: 'Please input your city!',
                  },
                ]}
                options={[
                  {
                    label: 'Gainesville',
                    value: 'Gainesville',
                  },
                  {
                    label: 'New York',
                    value: 'New York',
                  },
                  {
                    labe: 'Atalanta',
                    value: 'Atalanta',
                  },
                ]}
                initialValue={currentUser.city}
              />    
              
              <ProFormText
                width="md"
                name="phone"
                label="phone"
                rules={[
                  {
                    required: true,
                    message: 'Please input your phone number!',
                  },
                ]}
                initialValue={currentUser.phone}
              />

              {/* <ProFormFieldSet
                name="phone"
                label="phone"
                rules={[
                  {
                    required: true,
                    message: 'please input your phone!',
                  },
                  {
                    validator: validatorPhone,
                  },
                ]}
              >
                <Input className={styles.area_code} />
                <Input className={styles.phone_number} />
              </ProFormFieldSet> */}
            </ProForm>
          </div>
          <div className={styles.right}>
            <AvatarView avatar={currentUser.avatar} />
          </div>
        </>
      )}
    </div>
  );
};

export default BaseView;
