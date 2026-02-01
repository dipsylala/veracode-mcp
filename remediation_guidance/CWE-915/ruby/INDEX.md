# CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes - Ruby

## LLM Guidance

Mass assignment vulnerabilities in Ruby on Rails occur when Rails automatically assigns request parameters to model attributes, allowing attackers to modify security-critical fields like `is_admin`, `role`, or `balance`. Always use Strong Parameters to allowlist permitted attributes, never use `update` with unfiltered params, and validate all input.

## Key Principles

- **Allowlist only safe attributes** - Never permit all parameters; explicitly define permitted fields
- **Separate create/update permissions** - Different actions may require different permitted attributes
- **Protect administrative fields** - Never permit `role`, `is_admin`, `user_id`, or similar security fields
- **Validate business logic** - Strong Parameters prevents mass assignment but doesn't validate values
- **Avoid legacy patterns** - Never use `attr_accessible` or permit all with `params.permit!`

## Remediation Steps

- Define private `*_params` methods in controllers using `params.require().permit()`
- Replace all `Model.new(params[ -model])` with `Model.new(model_params)`
- Replace `@model.update(params[ -model])` with `@model.update(model_params)`
- Review permitted attributes - remove any administrative or security-critical fields
- Add server-side validation for business rules and constraints
- Test by attempting to inject unauthorized parameters in requests

## Safe Pattern

```ruby
class UsersController < ApplicationController
  def update
    @user = User.find(params[:id])
    if @user.update(user_params)
      redirect_to @user
    else
      render :edit
    end
  end

  private
  def user_params
    params.require(:user).permit(:name, :email, :bio)
    # Never permit: :is_admin, :role, :account_balance
  end
end
```
