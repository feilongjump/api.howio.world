<x-mail::message>
# Hello!

以下为您的验证码

> {{ $code }}

此验证码 5 分钟内有效。

Thanks,<br>
{{ config('app.name') }}
</x-mail::message>
