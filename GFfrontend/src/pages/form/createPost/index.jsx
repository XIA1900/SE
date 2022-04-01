import { Card, message } from 'antd';
import ProForm, {
  ProFormDateRangePicker,
  ProFormDependency,
  ProFormDigit,
  ProFormRadio,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-form';
import { history, useRequest, useModel, useIntl } from 'umi';
import { PageContainer } from '@ant-design/pro-layout';
import { createPost } from '@/services/create';
import styles from './style.less';

const groupName = history.location.search.substring(1);

const BasicForm = () => {

  const { initialState, setInitialState } = useModel('@@initialState');
  const { currentUser } = initialState;
  const intl = useIntl();

  const onFinish = async (values) => {
    const date = new Date();
    
    const result = await createPost({
      groupName: groupName,
      userName: currentUser.name,
      title: values.title,
      content: values.content,
      time: date.getFullYear()+"-"+date.getMonth()+"-"+date.getDate(),
    });

    console.log(result);


    if(result.message === 'Ok') {
      const defaultLoginSuccessMessage = intl.formatMessage({
        id: 'createPost',
        defaultMessage: 'Post submitted successfully!',
      });
      message.success(defaultLoginSuccessMessage);

      const postid = result.postid;

      history.push({
        pathname: '/group/post',
        search: postid,
      });
    }
    
  };

  return (
    <PageContainer content="What's in your mind?">
      <Card bordered={false}>
        <ProForm
          hideRequiredMark
          style={{
            margin: 'auto',
            marginTop: 8,
            maxWidth: 600,
          }}
          name="basic"
          layout="vertical"
          initialValues={{
            public: '1',
          }}
          onFinish={onFinish}
        >
          <ProFormText
            width="md"
            label="Title"
            name="title"
            rules={[
              {
                required: true,
                message: 'Please input a title for your post.',
              },
            ]}
            placeholder=""
          />

          <ProFormTextArea
            label="Content"
            width="xl"
            name="content"
            rules={[
              {
                required: true,
                message: 'Please add content for your post.',
              },
            ]}
            placeholder=""
          />
          </ProForm>
      </Card>
    </PageContainer>
  );
};

export default BasicForm;
