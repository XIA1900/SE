/*
updateGroupInfo, deleteGroup do not work
*/

import React, { useCallback, useRef } from 'react';
import { UploadOutlined } from '@ant-design/icons';
import { Form, Button, Input, Upload, message } from 'antd';
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
  const [form] = Form.useForm();

  const getAvatarURL = () => {
    if (list) {
      if (list.avatar) {
        return list.avatar;
      }
    }
    return '';
  };

  const onFinish = (values) => {
    console.log(values);
    const result = updateGroupInfo({
      values,
      
    });
    const msg = result.message;
    if(msg === 'Ok') {
      console.log("good");
    }
  }

  const onDelete = async () => {
    const msg = deleteGroup({groupName,}).msg;
      if(msg === 'Ok') {
        console.log('deleted');
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
            {/*begin change*/}
            <Form layout='horizontal' form={form} onFinish={onFinish}>

              <Form.Item 
                label="Group ID" 
                name='id' 
                initialValue={list.groupId} 
              >
                <Input disabled={true} />
              </Form.Item>

              <Form.Item 
                label="Group Owner" 
                name='owner' 
                initialValue={list.owner} 
              >
                <Input disabled={true} />
              </Form.Item>

              <Form.Item 
                label="Group Name" 
                name='name' 
                initialValue={list.name} 
                required
                tooltip='Please input a group name.'
              >
                <Input />
              </Form.Item>

              <Form.Item 
                label="Group Description" 
                name='description' 
                initialValue={list.description} 
                required
                tooltip='Please input group description.'
              >
                <Input.TextArea />
              </Form.Item>

              <Form.Item 
                label="Created At" 
                name='createdAt' 
                initialValue={list.createdAt} 
              >
                <Input disabled={true}/>
              </Form.Item>

              <Form.Item>
                <Button htmlType='submit' type='primary'>
                  Update
                </Button>
                <Button type='button' onClick={onDelete}>
                  Delete
                </Button>
              </Form.Item>
              
            </Form>
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


{/* <ProForm
              layout="vertical"
              submitter={false}
              // submitter={{
              //   render: (props, doms) => {
              //     return [
              //       //...doms,
              //       <Button htmlType="button" type="primary" onClick={(e) => handleUpdate(newGroupName, newDescription)} key="update">
              //         Update Information
              //       </Button>,
              //       <Button htmlType="button" danger onClick={handleDelete} key="delete">
              //         Delete Group
              //       </Button>,
              //     ];
              //   },
              // }}
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
              <ProForm.Item>
                <Button htmlType='button' onClick={handleUpdate}>
                  Update
                </Button>
                <Button htmlType='button' onClick={handleDelete}>
                  Delete
                </Button>
              </ProForm.Item>
              
            </ProForm> */}