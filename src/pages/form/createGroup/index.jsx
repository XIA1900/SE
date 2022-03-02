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
import { history, useRequest, useModel } from 'umi';
import { PageContainer } from '@ant-design/pro-layout';
import { createGroup } from '@/services/create';
import styles from './style.less';

const groupForm = () => {
  const { initialState, setInitialState } = useModel('@@initialState');
  const { currentUser } = initialState;
  console.log('currentUser');
  console.log(currentUser);
  const onFinish = async (values) => {
    
    let newdate = new Date();
    var date, month;
    if(newdate.getDate() < 10) 
      date = '0'+newdate.getDate().toString();
    else 
      date = newdate.getDate().toString();
    if(newdate.getMonth()+1<10) 
      month = '0'+(newdate.getMonth()+1).toString();
    else 
      month = (newdate.getMonth()+1).toString();
    let year = newdate.getFullYear();

    const params = {
      groupName: values.title,
      groupDescription: values.content,
      time: year+'-'+month+'-'+date,
      userId: currentUser.userId,
    };
    console.log(params);

    try {
      const msg = await createGroup({...params});
      if(msg.message == 'Ok') {
        history.push('/result/success');
      }
      else {

      }

    }
    catch(error){
      message.error("error");
    }
  };

  return (
    <PageContainer content="">
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
            label="Group Name"
            name="title"
            rules={[
              {
                required: true,
                message: 'Please enter a name for your group',
              },
            ]}
            placeholder="Please enter a name for your group"
          />

          <ProFormTextArea
            label="Group Description"
            width="xl"
            name="content"
            rules={[
              {
                required: true,
                message: 'Please enter a description for your group',
              },
            ]}
            placeholder="Please enter a description for your group"
          />
        </ProForm>
      </Card>
    </PageContainer>
  );
};

export default groupForm;
