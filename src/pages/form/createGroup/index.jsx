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
import { history, useRequest } from 'umi';
import { PageContainer } from '@ant-design/pro-layout';
import { createGroup } from '@/services/create';
import styles from './style.less';

const groupForm = () => {
  const { run } = useRequest(createGroup, {
    
  });

  const onFinish = async (values) => {
    run(values);
  };

  return (
    <PageContainer content="Please enter group information to create a group.">
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
                message: 'Please enter a name for your group',
              },
            ]}
            placeholder=""
          />

          <ProFormTextArea
            label="Content"
            width="xl"
            name="goal"
            rules={[
              {
                required: true,
                message: 'Please enter the description for your group',
              },
            ]}
            placeholder=""
          />
        </ProForm>
      </Card>
    </PageContainer>
  );
};

export default groupForm;
