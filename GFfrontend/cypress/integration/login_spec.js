describe('login', function(){
    this.beforeEach(()=>{
            cy.visit('http://localhost:8001/user/login')  // specify full URL if baseUrl is null or the domain is different the baseUrl
        })

    it('correctLogin', function(){
    cy.get('.ant-pro-form-login-main').get('#username').type('kirby')
    cy.get('.ant-pro-form-login-main').get('#password').type('007')
    cy.get('button:contains("Login")').click()
    cy.url().should('include', '/homepage')
    })
    it('incorrectLogin', function(){
    cy.get('.ant-pro-form-login-main').get('#username').type('admin')
    cy.get('.ant-pro-form-login-main').get('#password').type('ant')
    cy.get('button:contains("Login")').click()
    cy.url().should('include', '/user/login') 
    })
})
