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
import { useRequest } from 'umi';
import { PageContainer } from '@ant-design/pro-layout';
import { fakeSubmitForm } from './service';
import styles from './style.less';

const BasicForm = () => {
  const { run } = useRequest(fakeSubmitForm, {
    manual: true,
    onSuccess: () => {
      message.success('提交成功');
    },
  });

  const onFinish = async (values) => {
    run(values);
  };

  return (
    <PageContainer content="表单页用于向用户收集或验证信息，基础表单常见于数据项较少的表单场景。">
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
            label="title"
            name="title"
            rules={[
              {
                required: true,
                message: '请输入标题',
              },
            ]}
            placeholder=""
          />
          
          <ProFormTextArea
            label="content"
            width="xl"
            name="goal"
            rules={[
              {
                required: true,
                message: '请输入目标描述',
              },
            ]}
            placeholder=""
          />
          <ProFormRadio.Group
            options={[
              {
                value: '1',
                label: 'public',
              },
              {
                value: '2',
                label: 'private',
              },
            ]}
            label="share with"
            help=""
            name="publicType"
          />
        </ProForm>
      </Card>
    </PageContainer>
  );
};

export default BasicForm;
