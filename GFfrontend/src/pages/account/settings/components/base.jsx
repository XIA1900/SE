import React from 'react';
import { UploadOutlined } from '@ant-design/icons';
import { Button, Input, Upload, message } from 'antd';
import ProForm, {
  ProFormDependency,
  ProFormFieldSet,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-form';
import { useRequest } from 'umi';
import { queryCurrent } from '@/services/user';
import styles from './BaseView.less';

const validatorPhone = (rule, value, callback) => {
  if (!value[0]) {
    callback('Please input your area code!');
  }

  if (!value[1]) {
    callback('Please input your phone number!');
  }

  callback();
}; // 头像组件 方便以后独立，增加裁剪之类的功能

const AvatarView = ({ avatar }) => (
  <>
    <div className={styles.avatar_title}></div>
    <div className={styles.avatar}>
      <img src={avatar} alt="avatar" />
    </div>
    <Upload showUploadList={false}>
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
  const { data: currentUser, loading } = useRequest(() => {
    return queryCurrent();
  });

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

  const handleFinish = async () => {
    message.success('Change basic information successfully');
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
              initialValues={{ ...currentUser, phone: currentUser?.phone.split('-') }}
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
            <AvatarView avatar={getAvatarURL()} />
          </div>
        </>
      )}
    </div>
  );
};

export default BaseView;
