import React, { useRef } from 'react';
import { UploadOutlined } from '@ant-design/icons';
import { Button, Input, Upload, message } from 'antd';
import ProForm, {
  ProFormDependency,
  ProFormFieldSet,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-form';
import { ProFormInstance } from '@ant-design/pro-form';
import { useRequest, history } from 'umi';
import { getBasicInfo, updateGroupInfo, deleteGroup } from '@/services/groupManagement';
import styles from './BaseView.less';

const groupName = history.location.search.substring(1);
const formRef = React.createRef();
// 头像组件 方便以后独立，增加裁剪之类的功能
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

const BasicInfo = () => {
  const { data: basicInfo, loading } = useRequest(() => {
    return getBasicInfo({
        groupName,
    });
  });
  const list = basicInfo?.list || [];
  

  const getAvatarURL = () => {
    if (list) {
      if (list.avatar) {
        return list.avatar;
      }
    }
    return '';
  };

  const handleUpdate = async () => {
      console.log(this.formRef.current);
      const msg = "Ok";
      //const msg = updateGroupInfo({values,}).msg;
      if(msg === 'Ok') {

      }
      else if(msg === '') {

      }
      else {

      }
  };

  const handleDelete = async () => {
    const msg = deleteGroup({groupName,}).msg;
      if(msg === 'Ok') {

      }
      else if(msg === '') {

      }
      else {

      }
  };

  return (
    <div className={styles.baseView}>
      {loading ? null : (
        <>
          <div className={styles.left}>
            <ProForm
              useRef={formRef}
              layout="vertical"
              submitter={{
                render: (props, doms) => {
                  return [
                    //...doms,
                    <Button htmlType="button" type="primary" onClick={(e) => handleUpdate(newGroupName, newDescription)} key="update">
                      Update Information
                    </Button>,
                    <Button htmlType="button" danger onClick={handleDelete} key="delete">
                      Delete Group
                    </Button>,
                  ];
                },
              }}
              initialValues={{}}
              //initialValues={{ ...currentUser, phone: currentUser?.phone.split('-') }}
              hideRequiredMark
            >
              <ProFormText
                width="md"
                name="id"
                disabled label="Group ID"
                rules={[
                  {
                    required: true,
                    message: '',
                  },
                ]}
                initialValue={list.groupId}
              />
              <ProFormText
                width="md"
                name="owner"
                disabled label="Group Owner"
                rules={[
                  {
                    required: true,
                    message: '',
                  },
                ]}
                initialValue={list.owner}
              />
              <ProFormText
                width="md"
                name="name"
                label="Group Name"
                rules={[
                  {
                    required: true,
                    message: 'Please input a group name!',
                  },
                ]}
                initialValue = {list.name}
              />
              <ProFormTextArea
                name="description"
                label="Description"
                rules={[
                  {
                    required: true,
                    message: 'Please input group description!',
                  },
                ]}
                initialValue={list.description}
              />
              <ProFormText
                width="sm"
                name="createdAt"
                disabled label="Created At"
                rules={[
                  {
                    required: true,
                    message: '',
                  },
                ]}
                initialValue={list.createdAt}
              />
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

export default BasicInfo;
